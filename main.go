package main

import (
	cmd "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/cmd"
	_ "github.com/rookie-ninja/rk-gin/v2/boot"
)

// @title           Bitwyre P2P Develop
// @version         1.0
// @description     This is P2P Gateway Documentation.

func main() {
	cmd.Execute()
}
