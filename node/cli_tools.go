package node

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FlagsToArgs converts a `Flags` struct into a CLI command flag string
func FlagsToArgs(flags Flags, basedir string, sepBase bool) ([]string, Metadata) {
	// Port targets
	httpPortString := strconv.FormatUint(uint64(flags.HTTPPort), 10)
	stakingPortString := strconv.FormatUint(uint64(flags.StakingPort), 10)

	// Paths/directories
	dbPath := basedir + "/" + flags.DBDir
	dataPath := basedir + "/" + flags.DataDir
	logPath := basedir + "/" + flags.LogDir
	if sepBase {
		dbPath = flags.DBDir
		dataPath = flags.DataDir
		logPath = flags.LogDir
	}

	wd, _ := os.Getwd()
	// If the path given in the flag doesn't begin with "/", treat it as relative
	// to the directory of the avash binary
	httpCertFile := flags.HTTPTLSCertFile
	if httpCertFile != "" && string(httpCertFile[0]) != "/" && !sepBase {
		httpCertFile = fmt.Sprintf("%s/%s", wd, httpCertFile)
	}

	httpKeyFile := flags.HTTPTLSKeyFile
	if httpKeyFile != "" && string(httpKeyFile[0]) != "/" && !sepBase {
		httpKeyFile = fmt.Sprintf("%s/%s", wd, httpKeyFile)
	}

	stakerCertFile := flags.StakingTLSCertFile
	if stakerCertFile != "" && string(stakerCertFile[0]) != "/" && !sepBase {
		stakerCertFile = fmt.Sprintf("%s/%s", wd, stakerCertFile)
	}

	stakerKeyFile := flags.StakingTLSKeyFile
	if stakerKeyFile != "" && string(stakerKeyFile[0]) != "/" && !sepBase {
		stakerKeyFile = fmt.Sprintf("%s/%s", wd, stakerKeyFile)
	}

	args := []string{
		"--assertions-enabled=" + strconv.FormatBool(flags.AssertionsEnabled),
		"--tx-fee=" + strconv.FormatUint(uint64(flags.TxFee), 10),
		"--public-ip=" + flags.PublicIP,
		"--network-id=" + flags.NetworkID,
		"--xput-server-port=" + strconv.FormatUint(uint64(flags.XputServerPort), 10),
		"--xput-server-enabled=" + strconv.FormatBool(flags.XputServerEnabled),
		"--signature-verification-enabled=" + strconv.FormatBool(flags.SignatureVerificationEnabled),
		"--api-admin-enabled=" + strconv.FormatBool(flags.APIAdminEnabled),
		"--api-ipcs-enabled=" + strconv.FormatBool(flags.APIIPCsEnabled),
		"--api-keystore-enabled=" + strconv.FormatBool(flags.APIKeystoreEnabled),
		"--api-metrics-enabled=" + strconv.FormatBool(flags.APIMetricsEnabled),
		"--http-host=" + flags.HTTPHost,
		"--http-port=" + httpPortString,
		"--http-tls-enabled=" + strconv.FormatBool(flags.HTTPTLSEnabled),
		"--http-tls-cert-file=" + httpCertFile,
		"--http-tls-key-file=" + httpKeyFile,
		"--bootstrap-ips=" + flags.BootstrapIPs,
		"--bootstrap-ids=" + flags.BootstrapIDs,
		"--db-enabled=" + strconv.FormatBool(flags.DBEnabled),
		"--db-dir=" + dbPath,
		"--plugin-dir=" + flags.PluginDir,
		"--log-level=" + flags.LogLevel,
		"--log-dir=" + logPath,
		"--log-display-level=" + flags.LogDisplayLevel,
		"--snow-avalanche-batch-size=" + strconv.Itoa(flags.SnowAvalancheBatchSize),
		"--snow-avalanche-num-parents=" + strconv.Itoa(flags.SnowAvalancheNumParents),
		"--snow-sample-size=" + strconv.Itoa(flags.SnowSampleSize),
		"--snow-quorum-size=" + strconv.Itoa(flags.SnowQuorumSize),
		"--snow-virtuous-commit-threshold=" + strconv.Itoa(flags.SnowVirtuousCommitThreshold),
		"--snow-rogue-commit-threshold=" + strconv.Itoa(flags.SnowRogueCommitThreshold),
		"--p2p-tls-enabled=" + strconv.FormatBool(flags.P2PTLSEnabled),
		"--staking-enabled=" + strconv.FormatBool(flags.StakingEnabled),
		"--staking-port=" + stakingPortString,
		"--staking-tls-key-file=" + stakerKeyFile,
		"--staking-tls-cert-file=" + stakerCertFile,
		"--api-auth-required=" + strconv.FormatBool(flags.APIAuthRequired),
		"--api-auth-password=" + flags.APIAuthPassword,
		"--min-stake-duration=" + flags.MinStakeDuration,
		"--whitelisted-subnets=" + flags.WhitelistedSubnets,
		"--config-file=" + flags.ConfigFile,
		"--api-info-enabled=" + strconv.FormatBool(flags.APIInfoEnabled),
		"--conn-meter-max-conns=" + strconv.Itoa(flags.ConnMeterMaxConns),
		"--conn-meter-reset-duration=" + flags.ConnMeterResetDuration,
		"--ipcs-chain-ids=" + flags.IPCSChainIDs,
		"--ipcs-path=" + flags.IPCSPath,
		"--fd-limit=" + strconv.Itoa(flags.FDLimit),
	}
	if sepBase {
		args = append(args, "--data-dir="+basedir)
	}
	args = removeEmptyFlags(args)

	metadata := Metadata{
		Serverhost:     flags.PublicIP,
		Stakingport:    stakingPortString,
		HTTPport:       httpPortString,
		HTTPTLS:        flags.HTTPTLSEnabled,
		Dbdir:          dbPath,
		Datadir:        dataPath,
		Logsdir:        logPath,
		Loglevel:       flags.LogLevel,
		P2PTLSEnabled:  flags.P2PTLSEnabled,
		StakingEnabled: flags.StakingEnabled,
		StakerCertPath: stakerCertFile,
		StakerKeyPath:  stakerKeyFile,
	}

	return args, metadata
}

func removeEmptyFlags(args []string) []string {
	var res []string
	for _, f := range args {
		tmp := strings.TrimSpace(f)
		if !strings.HasSuffix(tmp, "=") {
			res = append(res, tmp)
		}
	}
	return res
}
