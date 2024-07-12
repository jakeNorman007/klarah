package flags

import (
    "fmt"
    "strings"
)

type Framework string

const (
    StandardLibrary Framework = "std-lib"
)

var FrameworkTypes = []string{string(StandardLibrary)}

func (f Framework) String() string {
    return string(f)
}

func (f *Framework) Type() string {
    return "Framework"
}

func (f *Framework) SetFramework(value string) error {
    for _, frmwrk := range FrameworkTypes {
        if frmwrk == value {
            *f = Framework(value)
            return nil
        }
    }

    return fmt.Errorf("Frameworks available to use: %s", strings.Join(FrameworkTypes, ", "))
}
