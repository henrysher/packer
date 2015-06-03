package interpolate

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/henrysher/packer/common/uuid"
)

// InitTime is the UTC time when this package was initialized. It is
// used as the timestamp for all configuration templates so that they
// match for a single build.
var InitTime time.Time

func init() {
	InitTime = time.Now().UTC()
}

// Funcs are the interpolation funcs that are available within interpolations.
var FuncGens = map[string]FuncGenerator{
	"env":       funcGenEnv,
	"isotime":   funcGenIsotime,
	"pwd":       funcGenPwd,
	"timestamp": funcGenTimestamp,
	"uuid":      funcGenUuid,
	"user":      funcGenUser,

	"upper": funcGenPrimitive(strings.ToUpper),
	"lower": funcGenPrimitive(strings.ToLower),
}

// FuncGenerator is a function that given a context generates a template
// function for the template.
type FuncGenerator func(*Context) interface{}

// Funcs returns the functions that can be used for interpolation given
// a context.
func Funcs(ctx *Context) template.FuncMap {
	result := make(map[string]interface{})
	for k, v := range FuncGens {
		result[k] = v(ctx)
	}

	return template.FuncMap(result)
}

func funcGenEnv(ctx *Context) interface{} {
	return func(k string) (string, error) {
		if !ctx.EnableEnv {
			// The error message doesn't have to be that detailed since
			// semantic checks should catch this.
			return "", errors.New("env vars are not allowed here")
		}

		return os.Getenv(k), nil
	}
}

func funcGenIsotime(ctx *Context) interface{} {
	return func(format ...string) (string, error) {
		if len(format) == 0 {
			return time.Now().UTC().Format(time.RFC3339), nil
		}

		if len(format) > 1 {
			return "", fmt.Errorf("too many values, 1 needed: %v", format)
		}

		return time.Now().UTC().Format(format[0]), nil
	}
}

func funcGenPrimitive(value interface{}) FuncGenerator {
	return func(ctx *Context) interface{} {
		return value
	}
}

func funcGenPwd(ctx *Context) interface{} {
	return func() (string, error) {
		return os.Getwd()
	}
}

func funcGenTimestamp(ctx *Context) interface{} {
	return func() string {
		return strconv.FormatInt(InitTime.Unix(), 10)
	}
}

func funcGenUser(ctx *Context) interface{} {
	return func(k string) string {
		if ctx == nil || ctx.UserVariables == nil {
			return ""
		}

		return ctx.UserVariables[k]
	}
}

func funcGenUuid(ctx *Context) interface{} {
	return func() string {
		return uuid.TimeOrderedUUID()
	}
}
