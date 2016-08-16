// Auto-generated by avdl-compiler v1.3.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/rekey.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type TLF struct {
	Id        TLFID    `codec:"id" json:"id"`
	Name      string   `codec:"name" json:"name"`
	Writers   []string `codec:"writers" json:"writers"`
	Readers   []string `codec:"readers" json:"readers"`
	IsPrivate bool     `codec:"isPrivate" json:"isPrivate"`
}

type ProblemTLF struct {
	Tlf           TLF   `codec:"tlf" json:"tlf"`
	Score         int   `codec:"score" json:"score"`
	Solution_kids []KID `codec:"solution_kids" json:"solution_kids"`
}

// ProblemSet is for a particular (user,kid) that initiated a rekey problem.
// This problem consists of one or more problem TLFs, which are individually scored
// and have attendant solutions --- devices that if they came online can rekey and
// solve the ProblemTLF.
type ProblemSet struct {
	User User         `codec:"user" json:"user"`
	Kid  KID          `codec:"kid" json:"kid"`
	Tlfs []ProblemTLF `codec:"tlfs" json:"tlfs"`
}

type ProblemSetDevices struct {
	ProblemSet ProblemSet `codec:"problemSet" json:"problemSet"`
	Devices    []Device   `codec:"devices" json:"devices"`
}

type Outcome int

const (
	Outcome_NONE    Outcome = 0
	Outcome_FIXED   Outcome = 1
	Outcome_IGNORED Outcome = 2
)

type ShowPendingRekeyStatusArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type GetPendingRekeyStatusArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type DebugShowRekeyStatusArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type RekeyStatusFinishArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type RekeySyncArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	Force     bool `codec:"force" json:"force"`
}

type RekeyInterface interface {
	// ShowPendingRekeyStatus shows either pending gregor-initiated rekey harassments
	// or nothing if none were pending.
	ShowPendingRekeyStatus(context.Context, int) error
	// GetPendingRekeyStatus returns the pending ProblemSetDevices.
	GetPendingRekeyStatus(context.Context, int) (ProblemSetDevices, error)
	// DebugShowRekeyStatus is used by the CLI to kick off a "ShowRekeyStatus" window for
	// the current user.
	DebugShowRekeyStatus(context.Context, int) error
	// RekeyStatusFinish is called when work is completed on a given RekeyStatus window. The Outcome
	// can be Fixed or Ignored.
	RekeyStatusFinish(context.Context, int) (Outcome, error)
	// RekeySync flushes the current rekey loop and gets to a good stopping point
	// to assert state. Good for race-free testing, not very useful in production.
	// Force overrides a long-snooze.
	RekeySync(context.Context, RekeySyncArg) error
}

func RekeyProtocol(i RekeyInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.rekey",
		Methods: map[string]rpc.ServeHandlerDescription{
			"showPendingRekeyStatus": {
				MakeArg: func() interface{} {
					ret := make([]ShowPendingRekeyStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ShowPendingRekeyStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]ShowPendingRekeyStatusArg)(nil), args)
						return
					}
					err = i.ShowPendingRekeyStatus(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getPendingRekeyStatus": {
				MakeArg: func() interface{} {
					ret := make([]GetPendingRekeyStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetPendingRekeyStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetPendingRekeyStatusArg)(nil), args)
						return
					}
					ret, err = i.GetPendingRekeyStatus(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"debugShowRekeyStatus": {
				MakeArg: func() interface{} {
					ret := make([]DebugShowRekeyStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DebugShowRekeyStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]DebugShowRekeyStatusArg)(nil), args)
						return
					}
					err = i.DebugShowRekeyStatus(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"rekeyStatusFinish": {
				MakeArg: func() interface{} {
					ret := make([]RekeyStatusFinishArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RekeyStatusFinishArg)
					if !ok {
						err = rpc.NewTypeError((*[]RekeyStatusFinishArg)(nil), args)
						return
					}
					ret, err = i.RekeyStatusFinish(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"rekeySync": {
				MakeArg: func() interface{} {
					ret := make([]RekeySyncArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RekeySyncArg)
					if !ok {
						err = rpc.NewTypeError((*[]RekeySyncArg)(nil), args)
						return
					}
					err = i.RekeySync(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type RekeyClient struct {
	Cli rpc.GenericClient
}

// ShowPendingRekeyStatus shows either pending gregor-initiated rekey harassments
// or nothing if none were pending.
func (c RekeyClient) ShowPendingRekeyStatus(ctx context.Context, sessionID int) (err error) {
	__arg := ShowPendingRekeyStatusArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.rekey.showPendingRekeyStatus", []interface{}{__arg}, nil)
	return
}

// GetPendingRekeyStatus returns the pending ProblemSetDevices.
func (c RekeyClient) GetPendingRekeyStatus(ctx context.Context, sessionID int) (res ProblemSetDevices, err error) {
	__arg := GetPendingRekeyStatusArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.rekey.getPendingRekeyStatus", []interface{}{__arg}, &res)
	return
}

// DebugShowRekeyStatus is used by the CLI to kick off a "ShowRekeyStatus" window for
// the current user.
func (c RekeyClient) DebugShowRekeyStatus(ctx context.Context, sessionID int) (err error) {
	__arg := DebugShowRekeyStatusArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.rekey.debugShowRekeyStatus", []interface{}{__arg}, nil)
	return
}

// RekeyStatusFinish is called when work is completed on a given RekeyStatus window. The Outcome
// can be Fixed or Ignored.
func (c RekeyClient) RekeyStatusFinish(ctx context.Context, sessionID int) (res Outcome, err error) {
	__arg := RekeyStatusFinishArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.rekey.rekeyStatusFinish", []interface{}{__arg}, &res)
	return
}

// RekeySync flushes the current rekey loop and gets to a good stopping point
// to assert state. Good for race-free testing, not very useful in production.
// Force overrides a long-snooze.
func (c RekeyClient) RekeySync(ctx context.Context, __arg RekeySyncArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.rekey.rekeySync", []interface{}{__arg}, nil)
	return
}
