// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package pack

import "kraftkit.sh/utils"

// PullOptions contains the list of options which can be set for pulling a
// package.
type PullOptions struct {
	architectures     []string
	platforms         []string
	version           string
	calculateChecksum bool
	onProgress        func(progress float64)
	workdir           string
	useCache          bool
}

// OnProgress calls (if set) an embedded progress function which can be used to
// update an external progress bar, for example.
func (ppo *PullOptions) OnProgress(progress float64) {
	if ppo.onProgress != nil {
		ppo.onProgress(progress)
	}
}

// Workdir returns the set working directory as part of the pull request
func (ppo *PullOptions) Workdir() string {
	return ppo.workdir
}

// Version returns
func (ppo *PullOptions) Version() string {
	return ppo.version
}

// CalculateChecksum returns whether the pull request should perform a check of
// the resource sum.
func (ppo *PullOptions) CalculateChecksum() bool {
	return ppo.calculateChecksum
}

// UseCache returns whether the pull should redirect to using a local cache if
// available.
func (ppo *PullOptions) UseCache() bool {
	return ppo.useCache
}

// PullOption is an option function which is used to modify PullOptions.
type PullOption func(opts *PullOptions) error

// NewPullOptions creates PullOptions
func NewPullOptions(opts ...PullOption) (*PullOptions, error) {
	options := &PullOptions{}

	for _, o := range opts {
		err := o(options)
		if err != nil {
			return nil, err
		}
	}

	return options, nil
}

// WithPullArchitecture requests a given architecture (if applicable)
func WithPullArchitecture(archs ...string) PullOption {
	return func(opts *PullOptions) error {
		for _, arch := range archs {
			if arch == "" {
				continue
			}

			if utils.Contains(opts.architectures, arch) {
				continue
			}

			opts.architectures = append(opts.architectures, archs...)
		}

		return nil
	}
}

// WithPullPlatform requests a given platform (if applicable).
func WithPullPlatform(plats ...string) PullOption {
	return func(opts *PullOptions) error {
		for _, plat := range plats {
			if plat == "" {
				continue
			}

			if utils.Contains(opts.platforms, plat) {
				continue
			}

			opts.platforms = append(opts.platforms, plats...)
		}

		return nil
	}
}

// WithPullProgressFunc set an optional progress function which is used as a
// callback during the transmission of the package and the host.
func WithPullProgressFunc(onProgress func(progress float64)) PullOption {
	return func(opts *PullOptions) error {
		opts.onProgress = onProgress
		return nil
	}
}

// WithPullWorkdir set the working directory context of the pull such that the
// resources of the package are placed there appropriately.
func WithPullWorkdir(workdir string) PullOption {
	return func(opts *PullOptions) error {
		opts.workdir = workdir
		return nil
	}
}

// WithPullChecksum to set whether to calculate and compare the checksum of the
// package.
func WithPullChecksum(calc bool) PullOption {
	return func(opts *PullOptions) error {
		opts.calculateChecksum = calc
		return nil
	}
}

// WithPullCache to set whether use cache if possible.
func WithPullCache(cache bool) PullOption {
	return func(opts *PullOptions) error {
		opts.useCache = cache
		return nil
	}
}

// WithPullVersion sets the version that should be pulled.
func WithPullVersion(version string) PullOption {
	return func(opts *PullOptions) error {
		opts.version = version
		return nil
	}
}
