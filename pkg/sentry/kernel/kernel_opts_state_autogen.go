// automatically generated by stateify.

//go:build go1.1
// +build go1.1

package kernel

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *SpecialOpts) StateTypeName() string {
	return "pkg/sentry/kernel.SpecialOpts"
}

func (s *SpecialOpts) StateFields() []string {
	return []string{}
}

func (s *SpecialOpts) beforeSave() {}

// +checklocksignore
func (s *SpecialOpts) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
}

func (s *SpecialOpts) afterLoad() {}

// +checklocksignore
func (s *SpecialOpts) StateLoad(stateSourceObject state.Source) {
}

func init() {
	state.Register((*SpecialOpts)(nil))
}