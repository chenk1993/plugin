// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package guess

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/33cn/plugin/plugin/dapp/guess/commands"
	"github.com/33cn/plugin/plugin/dapp/guess/executor"
	"github.com/33cn/plugin/plugin/dapp/guess/rpc"
	"github.com/33cn/plugin/plugin/dapp/guess/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.PokerBullX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.GuessCmd,
		RPC:      rpc.Init,
	})
}
