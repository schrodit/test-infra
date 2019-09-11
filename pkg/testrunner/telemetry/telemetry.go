// Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package telemetry

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path"
	"syscall"
)

type Telemetry struct {
	log logr.Logger

	path     string
	interval string
	watchCmd *exec.Cmd

	rawResults string
}

func New(log logr.Logger, path, interval string) (*Telemetry, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	return &Telemetry{
		log: log,

		path:     path,
		interval: interval,
	}, nil
}

// Start starts the telemetry measurement with a specific kubeconfig to watch all shoots
func (c *Telemetry) Start(kubeconfigPath, resultDir string, out bool) (string, error) {
	if _, err := os.Stat(kubeconfigPath); os.IsNotExist(err) {
		return "", err
	}
	c.rawResults = path.Join(resultDir, "results.csv")

	c.watchCmd = exec.Command(c.path, "--kubeconfig", kubeconfigPath, "--output", resultDir, "--interval", c.interval, "--analyse=false")
	if out {
		c.watchCmd.Stdout = os.Stdout
		c.watchCmd.Stderr = os.Stderr
	}
	if err := c.watchCmd.Start(); err != nil {
		return "", err
	}
	c.log.V(3).Info("telemetry-controller started", "pid", c.watchCmd.Process.Pid)
	return c.rawResults, nil
}

// StartForShoot starts the telemetry measurement with a kubeconfig for a specific shoot
func (c *Telemetry) StartForShoot(shootName, shootNamespace, kubeconfigPath, output string, out bool) (string, error) {
	if _, err := os.Stat(kubeconfigPath); os.IsNotExist(err) {
		return "", err
	}
	c.rawResults = fmt.Sprintf("%s/results.csv", output)

	c.watchCmd = exec.Command(c.path, "--kubeconfig", kubeconfigPath,
		"--shoot-name", shootName, "--shoot-namespace", shootNamespace, "--output", output, "--interval", c.interval, "--disable-analyse")
	if out {
		c.watchCmd.Stdout = os.Stdout
		c.watchCmd.Stderr = os.Stderr
	}
	if err := c.watchCmd.Start(); err != nil {
		return "", err
	}
	fmt.Printf("Controller started with PID %d\n", c.watchCmd.Process.Pid)
	return c.rawResults, nil
}

// StopAndAnalyze stops the telemetry measurement and generates a result summary
func (c *Telemetry) StopAndAnalyze(resultDir string, out bool) (string, error) {
	if err := c.Stop(); err != nil {
		return "", err
	}
	return c.Analyze(resultDir, out)
}

// Stop stops the measurement of the telemetry controller
func (c *Telemetry) Stop() error {
	if err := c.watchCmd.Process.Signal(syscall.SIGTERM); err != nil {
		return errors.Wrap(err, "unable to send sigterm signal")
	}
	if err := c.watchCmd.Wait(); err != nil {
		return errors.Wrap(err, "error while waiting for measurement command to stop")
	}
	return nil
}

// Analyze analyzes the previously measured values and returns the path to the summary
func (c *Telemetry) Analyze(resultDir string, out bool) (string, error) {
	summaryOutput := path.Join(resultDir, "summary.csv")

	cmd := exec.Command(c.path, "analyse", "--input", c.rawResults, "--report", summaryOutput, "--format", "json")
	if out {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	if err := cmd.Start(); err != nil {
		return "", errors.Wrap(err, "unable to start analyse command")
	}
	if err := cmd.Wait(); err != nil {
		return "", errors.Wrap(err, "error while waiting for analyse command to finish")
	}
	return summaryOutput, nil
}
