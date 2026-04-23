// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AppController

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IAppControllerAppConfig is an auto generated low-level Go binding around an user-defined struct.
type IAppControllerAppConfig struct {
	Owner                    common.Address
	OperatorSetId            uint32
	LatestReleaseBlockNumber uint32
	Status                   uint8
	Timelocked               bool
}

// IAppControllerAppRoles is an auto generated low-level Go binding around an user-defined struct.
type IAppControllerAppRoles struct {
	App     common.Address
	IsOwner bool
	Roles   []uint8
}

// IAppControllerRelease is an auto generated low-level Go binding around an user-defined struct.
type IAppControllerRelease struct {
	RmsRelease   IReleaseManagerTypesRelease
	PublicEnv    []byte
	EncryptedEnv []byte
}

// IReleaseManagerTypesArtifact is an auto generated low-level Go binding around an user-defined struct.
type IReleaseManagerTypesArtifact struct {
	Digest   [32]byte
	Registry string
}

// IReleaseManagerTypesRelease is an auto generated low-level Go binding around an user-defined struct.
type IReleaseManagerTypesRelease struct {
	Artifacts     []IReleaseManagerTypesArtifact
	UpgradeByTime uint32
}

// AppControllerMetaData contains all meta data concerning the AppController contract.
var AppControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_releaseManager\",\"type\":\"address\",\"internalType\":\"contractIReleaseManager\"},{\"name\":\"_computeAVSRegistrar\",\"type\":\"address\",\"internalType\":\"contractIComputeAVSRegistrar\"},{\"name\":\"_computeOperator\",\"type\":\"address\",\"internalType\":\"contractIComputeOperator\"},{\"name\":\"_appBeacon\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"},{\"name\":\"_safeTimelockFactory\",\"type\":\"address\",\"internalType\":\"contractISafeTimelockFactory\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"API_PERMISSION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"appBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateApiPermissionDigestHash\",\"inputs\":[{\"name\":\"permission\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateAppId\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canCall\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeAVSRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIComputeAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIComputeOperator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createApp\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIAppController.Release\",\"components\":[{\"name\":\"rmsRelease\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"publicEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encryptedEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAppForTeam\",\"inputs\":[{\"name\":\"team\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIAppController.Release\",\"components\":[{\"name\":\"rmsRelease\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"publicEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encryptedEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveAppCount\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppLatestReleaseBlockNumber\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppOperatorSetId\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppOwner\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppStatus\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.AppStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppTimelocked\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApps\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apps\",\"type\":\"address[]\",\"internalType\":\"contractIApp[]\"},{\"name\":\"appConfigsMem\",\"type\":\"tuple[]\",\"internalType\":\"structIAppController.AppConfig[]\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"latestReleaseBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.AppStatus\"},{\"name\":\"timelocked\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppsByCreator\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apps\",\"type\":\"address[]\",\"internalType\":\"contractIApp[]\"},{\"name\":\"appConfigsMem\",\"type\":\"tuple[]\",\"internalType\":\"structIAppController.AppConfig[]\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"latestReleaseBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.AppStatus\"},{\"name\":\"timelocked\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppsByDeveloper\",\"inputs\":[{\"name\":\"developer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apps\",\"type\":\"address[]\",\"internalType\":\"contractIApp[]\"},{\"name\":\"appConfigsMem\",\"type\":\"tuple[]\",\"internalType\":\"structIAppController.AppConfig[]\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"latestReleaseBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.AppStatus\"},{\"name\":\"timelocked\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppsForAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"appRoles\",\"type\":\"tuple[]\",\"internalType\":\"structIAppController.AppRoles[]\",\"components\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"},{\"name\":\"isOwner\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"roles\",\"type\":\"uint8[]\",\"internalType\":\"enumIAppController.TeamRole[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxActiveAppsPerUser\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTeamRoleMember\",\"inputs\":[{\"name\":\"team\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.TeamRole\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTeamRoleMemberCount\",\"inputs\":[{\"name\":\"team\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.TeamRole\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTeamRoleMembers\",\"inputs\":[{\"name\":\"team\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.TeamRole\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalActiveAppCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantTeamRole\",\"inputs\":[{\"name\":\"team\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.TeamRole\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasTeamRole\",\"inputs\":[{\"name\":\"team\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.TeamRole\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxGlobalActiveApps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrateAdmins\",\"inputs\":[{\"name\":\"apps\",\"type\":\"address[]\",\"internalType\":\"contractIApp[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIReleaseManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceTeamRole\",\"inputs\":[{\"name\":\"team\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.TeamRole\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeTeamRole\",\"inputs\":[{\"name\":\"team\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.TeamRole\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTimelockFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISafeTimelockFactory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMaxActiveAppsPerUser\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGlobalActiveApps\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startApp\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopApp\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"suspend\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"apps\",\"type\":\"address[]\",\"internalType\":\"contractIApp[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"terminateApp\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"terminateAppByAdmin\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAppMetadataURI\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeApp\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIAppController.Release\",\"components\":[{\"name\":\"rmsRelease\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"publicEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encryptedEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AppCreated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppMetadataURIUpdated\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppOwnershipTransferred\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"},{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppStarted\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppStopped\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppSuspended\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppTerminated\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppTerminatedByAdmin\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppUpgraded\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"},{\"name\":\"rmsReleaseId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"release\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAppController.Release\",\"components\":[{\"name\":\"rmsRelease\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"publicEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encryptedEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalMaxActiveAppsSet\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxActiveAppsSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccountHasActiveApps\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotRevokeLastAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalMaxActiveAppsExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAppStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReleaseMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxActiveAppsExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MoreThanOneArtifact\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// AppControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use AppControllerMetaData.ABI instead.
var AppControllerABI = AppControllerMetaData.ABI

// AppController is an auto generated Go binding around an Ethereum contract.
type AppController struct {
	AppControllerCaller     // Read-only binding to the contract
	AppControllerTransactor // Write-only binding to the contract
	AppControllerFilterer   // Log filterer for contract events
}

// AppControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppControllerSession struct {
	Contract     *AppController    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppControllerCallerSession struct {
	Contract *AppControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AppControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppControllerTransactorSession struct {
	Contract     *AppControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AppControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppControllerRaw struct {
	Contract *AppController // Generic contract binding to access the raw methods on
}

// AppControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppControllerCallerRaw struct {
	Contract *AppControllerCaller // Generic read-only contract binding to access the raw methods on
}

// AppControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppControllerTransactorRaw struct {
	Contract *AppControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppController creates a new instance of AppController, bound to a specific deployed contract.
func NewAppController(address common.Address, backend bind.ContractBackend) (*AppController, error) {
	contract, err := bindAppController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppController{AppControllerCaller: AppControllerCaller{contract: contract}, AppControllerTransactor: AppControllerTransactor{contract: contract}, AppControllerFilterer: AppControllerFilterer{contract: contract}}, nil
}

// NewAppControllerCaller creates a new read-only instance of AppController, bound to a specific deployed contract.
func NewAppControllerCaller(address common.Address, caller bind.ContractCaller) (*AppControllerCaller, error) {
	contract, err := bindAppController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppControllerCaller{contract: contract}, nil
}

// NewAppControllerTransactor creates a new write-only instance of AppController, bound to a specific deployed contract.
func NewAppControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*AppControllerTransactor, error) {
	contract, err := bindAppController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppControllerTransactor{contract: contract}, nil
}

// NewAppControllerFilterer creates a new log filterer instance of AppController, bound to a specific deployed contract.
func NewAppControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*AppControllerFilterer, error) {
	contract, err := bindAppController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppControllerFilterer{contract: contract}, nil
}

// bindAppController binds a generic wrapper to an already deployed contract.
func bindAppController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AppControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppController *AppControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppController.Contract.AppControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppController *AppControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppController.Contract.AppControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppController *AppControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppController.Contract.AppControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppController *AppControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppController *AppControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppController *AppControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppController.Contract.contract.Transact(opts, method, params...)
}

// APIPERMISSIONTYPEHASH is a free data retrieval call binding the contract method 0x8a7f922c.
//
// Solidity: function API_PERMISSION_TYPEHASH() view returns(bytes32)
func (_AppController *AppControllerCaller) APIPERMISSIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "API_PERMISSION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// APIPERMISSIONTYPEHASH is a free data retrieval call binding the contract method 0x8a7f922c.
//
// Solidity: function API_PERMISSION_TYPEHASH() view returns(bytes32)
func (_AppController *AppControllerSession) APIPERMISSIONTYPEHASH() ([32]byte, error) {
	return _AppController.Contract.APIPERMISSIONTYPEHASH(&_AppController.CallOpts)
}

// APIPERMISSIONTYPEHASH is a free data retrieval call binding the contract method 0x8a7f922c.
//
// Solidity: function API_PERMISSION_TYPEHASH() view returns(bytes32)
func (_AppController *AppControllerCallerSession) APIPERMISSIONTYPEHASH() ([32]byte, error) {
	return _AppController.Contract.APIPERMISSIONTYPEHASH(&_AppController.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AppController *AppControllerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AppController *AppControllerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AppController.Contract.DEFAULTADMINROLE(&_AppController.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AppController *AppControllerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AppController.Contract.DEFAULTADMINROLE(&_AppController.CallOpts)
}

// AppBeacon is a free data retrieval call binding the contract method 0x8a52d0b5.
//
// Solidity: function appBeacon() view returns(address)
func (_AppController *AppControllerCaller) AppBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "appBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppBeacon is a free data retrieval call binding the contract method 0x8a52d0b5.
//
// Solidity: function appBeacon() view returns(address)
func (_AppController *AppControllerSession) AppBeacon() (common.Address, error) {
	return _AppController.Contract.AppBeacon(&_AppController.CallOpts)
}

// AppBeacon is a free data retrieval call binding the contract method 0x8a52d0b5.
//
// Solidity: function appBeacon() view returns(address)
func (_AppController *AppControllerCallerSession) AppBeacon() (common.Address, error) {
	return _AppController.Contract.AppBeacon(&_AppController.CallOpts)
}

// CalculateApiPermissionDigestHash is a free data retrieval call binding the contract method 0x9690b7e7.
//
// Solidity: function calculateApiPermissionDigestHash(bytes4 permission, uint256 expiry) view returns(bytes32)
func (_AppController *AppControllerCaller) CalculateApiPermissionDigestHash(opts *bind.CallOpts, permission [4]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "calculateApiPermissionDigestHash", permission, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateApiPermissionDigestHash is a free data retrieval call binding the contract method 0x9690b7e7.
//
// Solidity: function calculateApiPermissionDigestHash(bytes4 permission, uint256 expiry) view returns(bytes32)
func (_AppController *AppControllerSession) CalculateApiPermissionDigestHash(permission [4]byte, expiry *big.Int) ([32]byte, error) {
	return _AppController.Contract.CalculateApiPermissionDigestHash(&_AppController.CallOpts, permission, expiry)
}

// CalculateApiPermissionDigestHash is a free data retrieval call binding the contract method 0x9690b7e7.
//
// Solidity: function calculateApiPermissionDigestHash(bytes4 permission, uint256 expiry) view returns(bytes32)
func (_AppController *AppControllerCallerSession) CalculateApiPermissionDigestHash(permission [4]byte, expiry *big.Int) ([32]byte, error) {
	return _AppController.Contract.CalculateApiPermissionDigestHash(&_AppController.CallOpts, permission, expiry)
}

// CalculateAppId is a free data retrieval call binding the contract method 0x36b18874.
//
// Solidity: function calculateAppId(address deployer, bytes32 salt) view returns(address)
func (_AppController *AppControllerCaller) CalculateAppId(opts *bind.CallOpts, deployer common.Address, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "calculateAppId", deployer, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateAppId is a free data retrieval call binding the contract method 0x36b18874.
//
// Solidity: function calculateAppId(address deployer, bytes32 salt) view returns(address)
func (_AppController *AppControllerSession) CalculateAppId(deployer common.Address, salt [32]byte) (common.Address, error) {
	return _AppController.Contract.CalculateAppId(&_AppController.CallOpts, deployer, salt)
}

// CalculateAppId is a free data retrieval call binding the contract method 0x36b18874.
//
// Solidity: function calculateAppId(address deployer, bytes32 salt) view returns(address)
func (_AppController *AppControllerCallerSession) CalculateAppId(deployer common.Address, salt [32]byte) (common.Address, error) {
	return _AppController.Contract.CalculateAppId(&_AppController.CallOpts, deployer, salt)
}

// CanCall is a free data retrieval call binding the contract method 0x9614801b.
//
// Solidity: function canCall(address caller, bytes data) view returns(bool)
func (_AppController *AppControllerCaller) CanCall(opts *bind.CallOpts, caller common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "canCall", caller, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCall is a free data retrieval call binding the contract method 0x9614801b.
//
// Solidity: function canCall(address caller, bytes data) view returns(bool)
func (_AppController *AppControllerSession) CanCall(caller common.Address, data []byte) (bool, error) {
	return _AppController.Contract.CanCall(&_AppController.CallOpts, caller, data)
}

// CanCall is a free data retrieval call binding the contract method 0x9614801b.
//
// Solidity: function canCall(address caller, bytes data) view returns(bool)
func (_AppController *AppControllerCallerSession) CanCall(caller common.Address, data []byte) (bool, error) {
	return _AppController.Contract.CanCall(&_AppController.CallOpts, caller, data)
}

// ComputeAVSRegistrar is a free data retrieval call binding the contract method 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (_AppController *AppControllerCaller) ComputeAVSRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "computeAVSRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeAVSRegistrar is a free data retrieval call binding the contract method 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (_AppController *AppControllerSession) ComputeAVSRegistrar() (common.Address, error) {
	return _AppController.Contract.ComputeAVSRegistrar(&_AppController.CallOpts)
}

// ComputeAVSRegistrar is a free data retrieval call binding the contract method 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (_AppController *AppControllerCallerSession) ComputeAVSRegistrar() (common.Address, error) {
	return _AppController.Contract.ComputeAVSRegistrar(&_AppController.CallOpts)
}

// ComputeOperator is a free data retrieval call binding the contract method 0x3fc4cbfd.
//
// Solidity: function computeOperator() view returns(address)
func (_AppController *AppControllerCaller) ComputeOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "computeOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeOperator is a free data retrieval call binding the contract method 0x3fc4cbfd.
//
// Solidity: function computeOperator() view returns(address)
func (_AppController *AppControllerSession) ComputeOperator() (common.Address, error) {
	return _AppController.Contract.ComputeOperator(&_AppController.CallOpts)
}

// ComputeOperator is a free data retrieval call binding the contract method 0x3fc4cbfd.
//
// Solidity: function computeOperator() view returns(address)
func (_AppController *AppControllerCallerSession) ComputeOperator() (common.Address, error) {
	return _AppController.Contract.ComputeOperator(&_AppController.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AppController *AppControllerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AppController *AppControllerSession) DomainSeparator() ([32]byte, error) {
	return _AppController.Contract.DomainSeparator(&_AppController.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AppController *AppControllerCallerSession) DomainSeparator() ([32]byte, error) {
	return _AppController.Contract.DomainSeparator(&_AppController.CallOpts)
}

// GetActiveAppCount is a free data retrieval call binding the contract method 0x0c2199fb.
//
// Solidity: function getActiveAppCount(address user) view returns(uint32)
func (_AppController *AppControllerCaller) GetActiveAppCount(opts *bind.CallOpts, user common.Address) (uint32, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getActiveAppCount", user)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetActiveAppCount is a free data retrieval call binding the contract method 0x0c2199fb.
//
// Solidity: function getActiveAppCount(address user) view returns(uint32)
func (_AppController *AppControllerSession) GetActiveAppCount(user common.Address) (uint32, error) {
	return _AppController.Contract.GetActiveAppCount(&_AppController.CallOpts, user)
}

// GetActiveAppCount is a free data retrieval call binding the contract method 0x0c2199fb.
//
// Solidity: function getActiveAppCount(address user) view returns(uint32)
func (_AppController *AppControllerCallerSession) GetActiveAppCount(user common.Address) (uint32, error) {
	return _AppController.Contract.GetActiveAppCount(&_AppController.CallOpts, user)
}

// GetAppLatestReleaseBlockNumber is a free data retrieval call binding the contract method 0x9ffbdce6.
//
// Solidity: function getAppLatestReleaseBlockNumber(address app) view returns(uint32)
func (_AppController *AppControllerCaller) GetAppLatestReleaseBlockNumber(opts *bind.CallOpts, app common.Address) (uint32, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getAppLatestReleaseBlockNumber", app)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAppLatestReleaseBlockNumber is a free data retrieval call binding the contract method 0x9ffbdce6.
//
// Solidity: function getAppLatestReleaseBlockNumber(address app) view returns(uint32)
func (_AppController *AppControllerSession) GetAppLatestReleaseBlockNumber(app common.Address) (uint32, error) {
	return _AppController.Contract.GetAppLatestReleaseBlockNumber(&_AppController.CallOpts, app)
}

// GetAppLatestReleaseBlockNumber is a free data retrieval call binding the contract method 0x9ffbdce6.
//
// Solidity: function getAppLatestReleaseBlockNumber(address app) view returns(uint32)
func (_AppController *AppControllerCallerSession) GetAppLatestReleaseBlockNumber(app common.Address) (uint32, error) {
	return _AppController.Contract.GetAppLatestReleaseBlockNumber(&_AppController.CallOpts, app)
}

// GetAppOperatorSetId is a free data retrieval call binding the contract method 0x6eb2099f.
//
// Solidity: function getAppOperatorSetId(address app) view returns(uint32)
func (_AppController *AppControllerCaller) GetAppOperatorSetId(opts *bind.CallOpts, app common.Address) (uint32, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getAppOperatorSetId", app)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAppOperatorSetId is a free data retrieval call binding the contract method 0x6eb2099f.
//
// Solidity: function getAppOperatorSetId(address app) view returns(uint32)
func (_AppController *AppControllerSession) GetAppOperatorSetId(app common.Address) (uint32, error) {
	return _AppController.Contract.GetAppOperatorSetId(&_AppController.CallOpts, app)
}

// GetAppOperatorSetId is a free data retrieval call binding the contract method 0x6eb2099f.
//
// Solidity: function getAppOperatorSetId(address app) view returns(uint32)
func (_AppController *AppControllerCallerSession) GetAppOperatorSetId(app common.Address) (uint32, error) {
	return _AppController.Contract.GetAppOperatorSetId(&_AppController.CallOpts, app)
}

// GetAppOwner is a free data retrieval call binding the contract method 0xffb42b51.
//
// Solidity: function getAppOwner(address app) view returns(address)
func (_AppController *AppControllerCaller) GetAppOwner(opts *bind.CallOpts, app common.Address) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getAppOwner", app)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAppOwner is a free data retrieval call binding the contract method 0xffb42b51.
//
// Solidity: function getAppOwner(address app) view returns(address)
func (_AppController *AppControllerSession) GetAppOwner(app common.Address) (common.Address, error) {
	return _AppController.Contract.GetAppOwner(&_AppController.CallOpts, app)
}

// GetAppOwner is a free data retrieval call binding the contract method 0xffb42b51.
//
// Solidity: function getAppOwner(address app) view returns(address)
func (_AppController *AppControllerCallerSession) GetAppOwner(app common.Address) (common.Address, error) {
	return _AppController.Contract.GetAppOwner(&_AppController.CallOpts, app)
}

// GetAppStatus is a free data retrieval call binding the contract method 0xd5aae178.
//
// Solidity: function getAppStatus(address app) view returns(uint8)
func (_AppController *AppControllerCaller) GetAppStatus(opts *bind.CallOpts, app common.Address) (uint8, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getAppStatus", app)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAppStatus is a free data retrieval call binding the contract method 0xd5aae178.
//
// Solidity: function getAppStatus(address app) view returns(uint8)
func (_AppController *AppControllerSession) GetAppStatus(app common.Address) (uint8, error) {
	return _AppController.Contract.GetAppStatus(&_AppController.CallOpts, app)
}

// GetAppStatus is a free data retrieval call binding the contract method 0xd5aae178.
//
// Solidity: function getAppStatus(address app) view returns(uint8)
func (_AppController *AppControllerCallerSession) GetAppStatus(app common.Address) (uint8, error) {
	return _AppController.Contract.GetAppStatus(&_AppController.CallOpts, app)
}

// GetAppTimelocked is a free data retrieval call binding the contract method 0x508ac1d6.
//
// Solidity: function getAppTimelocked(address app) view returns(bool)
func (_AppController *AppControllerCaller) GetAppTimelocked(opts *bind.CallOpts, app common.Address) (bool, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getAppTimelocked", app)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAppTimelocked is a free data retrieval call binding the contract method 0x508ac1d6.
//
// Solidity: function getAppTimelocked(address app) view returns(bool)
func (_AppController *AppControllerSession) GetAppTimelocked(app common.Address) (bool, error) {
	return _AppController.Contract.GetAppTimelocked(&_AppController.CallOpts, app)
}

// GetAppTimelocked is a free data retrieval call binding the contract method 0x508ac1d6.
//
// Solidity: function getAppTimelocked(address app) view returns(bool)
func (_AppController *AppControllerCallerSession) GetAppTimelocked(app common.Address) (bool, error) {
	return _AppController.Contract.GetAppTimelocked(&_AppController.CallOpts, app)
}

// GetApps is a free data retrieval call binding the contract method 0xa37e7e44.
//
// Solidity: function getApps(uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8,bool)[] appConfigsMem)
func (_AppController *AppControllerCaller) GetApps(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getApps", offset, limit)

	outstruct := new(struct {
		Apps          []common.Address
		AppConfigsMem []IAppControllerAppConfig
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Apps = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AppConfigsMem = *abi.ConvertType(out[1], new([]IAppControllerAppConfig)).(*[]IAppControllerAppConfig)

	return *outstruct, err

}

// GetApps is a free data retrieval call binding the contract method 0xa37e7e44.
//
// Solidity: function getApps(uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8,bool)[] appConfigsMem)
func (_AppController *AppControllerSession) GetApps(offset *big.Int, limit *big.Int) (struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}, error) {
	return _AppController.Contract.GetApps(&_AppController.CallOpts, offset, limit)
}

// GetApps is a free data retrieval call binding the contract method 0xa37e7e44.
//
// Solidity: function getApps(uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8,bool)[] appConfigsMem)
func (_AppController *AppControllerCallerSession) GetApps(offset *big.Int, limit *big.Int) (struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}, error) {
	return _AppController.Contract.GetApps(&_AppController.CallOpts, offset, limit)
}

// GetAppsByCreator is a free data retrieval call binding the contract method 0x8099ef2e.
//
// Solidity: function getAppsByCreator(address creator, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8,bool)[] appConfigsMem)
func (_AppController *AppControllerCaller) GetAppsByCreator(opts *bind.CallOpts, creator common.Address, offset *big.Int, limit *big.Int) (struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getAppsByCreator", creator, offset, limit)

	outstruct := new(struct {
		Apps          []common.Address
		AppConfigsMem []IAppControllerAppConfig
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Apps = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AppConfigsMem = *abi.ConvertType(out[1], new([]IAppControllerAppConfig)).(*[]IAppControllerAppConfig)

	return *outstruct, err

}

// GetAppsByCreator is a free data retrieval call binding the contract method 0x8099ef2e.
//
// Solidity: function getAppsByCreator(address creator, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8,bool)[] appConfigsMem)
func (_AppController *AppControllerSession) GetAppsByCreator(creator common.Address, offset *big.Int, limit *big.Int) (struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}, error) {
	return _AppController.Contract.GetAppsByCreator(&_AppController.CallOpts, creator, offset, limit)
}

// GetAppsByCreator is a free data retrieval call binding the contract method 0x8099ef2e.
//
// Solidity: function getAppsByCreator(address creator, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8,bool)[] appConfigsMem)
func (_AppController *AppControllerCallerSession) GetAppsByCreator(creator common.Address, offset *big.Int, limit *big.Int) (struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}, error) {
	return _AppController.Contract.GetAppsByCreator(&_AppController.CallOpts, creator, offset, limit)
}

// GetAppsByDeveloper is a free data retrieval call binding the contract method 0xf36618ac.
//
// Solidity: function getAppsByDeveloper(address developer, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8,bool)[] appConfigsMem)
func (_AppController *AppControllerCaller) GetAppsByDeveloper(opts *bind.CallOpts, developer common.Address, offset *big.Int, limit *big.Int) (struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getAppsByDeveloper", developer, offset, limit)

	outstruct := new(struct {
		Apps          []common.Address
		AppConfigsMem []IAppControllerAppConfig
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Apps = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AppConfigsMem = *abi.ConvertType(out[1], new([]IAppControllerAppConfig)).(*[]IAppControllerAppConfig)

	return *outstruct, err

}

// GetAppsByDeveloper is a free data retrieval call binding the contract method 0xf36618ac.
//
// Solidity: function getAppsByDeveloper(address developer, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8,bool)[] appConfigsMem)
func (_AppController *AppControllerSession) GetAppsByDeveloper(developer common.Address, offset *big.Int, limit *big.Int) (struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}, error) {
	return _AppController.Contract.GetAppsByDeveloper(&_AppController.CallOpts, developer, offset, limit)
}

// GetAppsByDeveloper is a free data retrieval call binding the contract method 0xf36618ac.
//
// Solidity: function getAppsByDeveloper(address developer, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8,bool)[] appConfigsMem)
func (_AppController *AppControllerCallerSession) GetAppsByDeveloper(developer common.Address, offset *big.Int, limit *big.Int) (struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}, error) {
	return _AppController.Contract.GetAppsByDeveloper(&_AppController.CallOpts, developer, offset, limit)
}

// GetAppsForAccount is a free data retrieval call binding the contract method 0xf61b3a25.
//
// Solidity: function getAppsForAccount(address account, uint256 offset, uint256 limit) view returns((address,bool,uint8[])[] appRoles)
func (_AppController *AppControllerCaller) GetAppsForAccount(opts *bind.CallOpts, account common.Address, offset *big.Int, limit *big.Int) ([]IAppControllerAppRoles, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getAppsForAccount", account, offset, limit)

	if err != nil {
		return *new([]IAppControllerAppRoles), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAppControllerAppRoles)).(*[]IAppControllerAppRoles)

	return out0, err

}

// GetAppsForAccount is a free data retrieval call binding the contract method 0xf61b3a25.
//
// Solidity: function getAppsForAccount(address account, uint256 offset, uint256 limit) view returns((address,bool,uint8[])[] appRoles)
func (_AppController *AppControllerSession) GetAppsForAccount(account common.Address, offset *big.Int, limit *big.Int) ([]IAppControllerAppRoles, error) {
	return _AppController.Contract.GetAppsForAccount(&_AppController.CallOpts, account, offset, limit)
}

// GetAppsForAccount is a free data retrieval call binding the contract method 0xf61b3a25.
//
// Solidity: function getAppsForAccount(address account, uint256 offset, uint256 limit) view returns((address,bool,uint8[])[] appRoles)
func (_AppController *AppControllerCallerSession) GetAppsForAccount(account common.Address, offset *big.Int, limit *big.Int) ([]IAppControllerAppRoles, error) {
	return _AppController.Contract.GetAppsForAccount(&_AppController.CallOpts, account, offset, limit)
}

// GetMaxActiveAppsPerUser is a free data retrieval call binding the contract method 0x40f70a9e.
//
// Solidity: function getMaxActiveAppsPerUser(address user) view returns(uint32)
func (_AppController *AppControllerCaller) GetMaxActiveAppsPerUser(opts *bind.CallOpts, user common.Address) (uint32, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getMaxActiveAppsPerUser", user)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMaxActiveAppsPerUser is a free data retrieval call binding the contract method 0x40f70a9e.
//
// Solidity: function getMaxActiveAppsPerUser(address user) view returns(uint32)
func (_AppController *AppControllerSession) GetMaxActiveAppsPerUser(user common.Address) (uint32, error) {
	return _AppController.Contract.GetMaxActiveAppsPerUser(&_AppController.CallOpts, user)
}

// GetMaxActiveAppsPerUser is a free data retrieval call binding the contract method 0x40f70a9e.
//
// Solidity: function getMaxActiveAppsPerUser(address user) view returns(uint32)
func (_AppController *AppControllerCallerSession) GetMaxActiveAppsPerUser(user common.Address) (uint32, error) {
	return _AppController.Contract.GetMaxActiveAppsPerUser(&_AppController.CallOpts, user)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AppController *AppControllerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AppController *AppControllerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AppController.Contract.GetRoleAdmin(&_AppController.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AppController *AppControllerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AppController.Contract.GetRoleAdmin(&_AppController.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AppController *AppControllerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AppController *AppControllerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AppController.Contract.GetRoleMember(&_AppController.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AppController *AppControllerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AppController.Contract.GetRoleMember(&_AppController.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AppController *AppControllerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AppController *AppControllerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AppController.Contract.GetRoleMemberCount(&_AppController.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AppController *AppControllerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AppController.Contract.GetRoleMemberCount(&_AppController.CallOpts, role)
}

// GetTeamRoleMember is a free data retrieval call binding the contract method 0x1d349e61.
//
// Solidity: function getTeamRoleMember(address team, uint8 role, uint256 index) view returns(address)
func (_AppController *AppControllerCaller) GetTeamRoleMember(opts *bind.CallOpts, team common.Address, role uint8, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getTeamRoleMember", team, role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeamRoleMember is a free data retrieval call binding the contract method 0x1d349e61.
//
// Solidity: function getTeamRoleMember(address team, uint8 role, uint256 index) view returns(address)
func (_AppController *AppControllerSession) GetTeamRoleMember(team common.Address, role uint8, index *big.Int) (common.Address, error) {
	return _AppController.Contract.GetTeamRoleMember(&_AppController.CallOpts, team, role, index)
}

// GetTeamRoleMember is a free data retrieval call binding the contract method 0x1d349e61.
//
// Solidity: function getTeamRoleMember(address team, uint8 role, uint256 index) view returns(address)
func (_AppController *AppControllerCallerSession) GetTeamRoleMember(team common.Address, role uint8, index *big.Int) (common.Address, error) {
	return _AppController.Contract.GetTeamRoleMember(&_AppController.CallOpts, team, role, index)
}

// GetTeamRoleMemberCount is a free data retrieval call binding the contract method 0x3c98ff65.
//
// Solidity: function getTeamRoleMemberCount(address team, uint8 role) view returns(uint256)
func (_AppController *AppControllerCaller) GetTeamRoleMemberCount(opts *bind.CallOpts, team common.Address, role uint8) (*big.Int, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getTeamRoleMemberCount", team, role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTeamRoleMemberCount is a free data retrieval call binding the contract method 0x3c98ff65.
//
// Solidity: function getTeamRoleMemberCount(address team, uint8 role) view returns(uint256)
func (_AppController *AppControllerSession) GetTeamRoleMemberCount(team common.Address, role uint8) (*big.Int, error) {
	return _AppController.Contract.GetTeamRoleMemberCount(&_AppController.CallOpts, team, role)
}

// GetTeamRoleMemberCount is a free data retrieval call binding the contract method 0x3c98ff65.
//
// Solidity: function getTeamRoleMemberCount(address team, uint8 role) view returns(uint256)
func (_AppController *AppControllerCallerSession) GetTeamRoleMemberCount(team common.Address, role uint8) (*big.Int, error) {
	return _AppController.Contract.GetTeamRoleMemberCount(&_AppController.CallOpts, team, role)
}

// GetTeamRoleMembers is a free data retrieval call binding the contract method 0x60257ae6.
//
// Solidity: function getTeamRoleMembers(address team, uint8 role) view returns(address[])
func (_AppController *AppControllerCaller) GetTeamRoleMembers(opts *bind.CallOpts, team common.Address, role uint8) ([]common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "getTeamRoleMembers", team, role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTeamRoleMembers is a free data retrieval call binding the contract method 0x60257ae6.
//
// Solidity: function getTeamRoleMembers(address team, uint8 role) view returns(address[])
func (_AppController *AppControllerSession) GetTeamRoleMembers(team common.Address, role uint8) ([]common.Address, error) {
	return _AppController.Contract.GetTeamRoleMembers(&_AppController.CallOpts, team, role)
}

// GetTeamRoleMembers is a free data retrieval call binding the contract method 0x60257ae6.
//
// Solidity: function getTeamRoleMembers(address team, uint8 role) view returns(address[])
func (_AppController *AppControllerCallerSession) GetTeamRoleMembers(team common.Address, role uint8) ([]common.Address, error) {
	return _AppController.Contract.GetTeamRoleMembers(&_AppController.CallOpts, team, role)
}

// GlobalActiveAppCount is a free data retrieval call binding the contract method 0xa8aa2bd3.
//
// Solidity: function globalActiveAppCount() view returns(uint32)
func (_AppController *AppControllerCaller) GlobalActiveAppCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "globalActiveAppCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GlobalActiveAppCount is a free data retrieval call binding the contract method 0xa8aa2bd3.
//
// Solidity: function globalActiveAppCount() view returns(uint32)
func (_AppController *AppControllerSession) GlobalActiveAppCount() (uint32, error) {
	return _AppController.Contract.GlobalActiveAppCount(&_AppController.CallOpts)
}

// GlobalActiveAppCount is a free data retrieval call binding the contract method 0xa8aa2bd3.
//
// Solidity: function globalActiveAppCount() view returns(uint32)
func (_AppController *AppControllerCallerSession) GlobalActiveAppCount() (uint32, error) {
	return _AppController.Contract.GlobalActiveAppCount(&_AppController.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AppController *AppControllerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AppController *AppControllerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AppController.Contract.HasRole(&_AppController.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AppController *AppControllerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AppController.Contract.HasRole(&_AppController.CallOpts, role, account)
}

// HasTeamRole is a free data retrieval call binding the contract method 0x54bfb170.
//
// Solidity: function hasTeamRole(address team, uint8 role, address account) view returns(bool)
func (_AppController *AppControllerCaller) HasTeamRole(opts *bind.CallOpts, team common.Address, role uint8, account common.Address) (bool, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "hasTeamRole", team, role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasTeamRole is a free data retrieval call binding the contract method 0x54bfb170.
//
// Solidity: function hasTeamRole(address team, uint8 role, address account) view returns(bool)
func (_AppController *AppControllerSession) HasTeamRole(team common.Address, role uint8, account common.Address) (bool, error) {
	return _AppController.Contract.HasTeamRole(&_AppController.CallOpts, team, role, account)
}

// HasTeamRole is a free data retrieval call binding the contract method 0x54bfb170.
//
// Solidity: function hasTeamRole(address team, uint8 role, address account) view returns(bool)
func (_AppController *AppControllerCallerSession) HasTeamRole(team common.Address, role uint8, account common.Address) (bool, error) {
	return _AppController.Contract.HasTeamRole(&_AppController.CallOpts, team, role, account)
}

// MaxGlobalActiveApps is a free data retrieval call binding the contract method 0xa46530a2.
//
// Solidity: function maxGlobalActiveApps() view returns(uint32)
func (_AppController *AppControllerCaller) MaxGlobalActiveApps(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "maxGlobalActiveApps")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxGlobalActiveApps is a free data retrieval call binding the contract method 0xa46530a2.
//
// Solidity: function maxGlobalActiveApps() view returns(uint32)
func (_AppController *AppControllerSession) MaxGlobalActiveApps() (uint32, error) {
	return _AppController.Contract.MaxGlobalActiveApps(&_AppController.CallOpts)
}

// MaxGlobalActiveApps is a free data retrieval call binding the contract method 0xa46530a2.
//
// Solidity: function maxGlobalActiveApps() view returns(uint32)
func (_AppController *AppControllerCallerSession) MaxGlobalActiveApps() (uint32, error) {
	return _AppController.Contract.MaxGlobalActiveApps(&_AppController.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_AppController *AppControllerCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_AppController *AppControllerSession) PermissionController() (common.Address, error) {
	return _AppController.Contract.PermissionController(&_AppController.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_AppController *AppControllerCallerSession) PermissionController() (common.Address, error) {
	return _AppController.Contract.PermissionController(&_AppController.CallOpts)
}

// ReleaseManager is a free data retrieval call binding the contract method 0x0d78573e.
//
// Solidity: function releaseManager() view returns(address)
func (_AppController *AppControllerCaller) ReleaseManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "releaseManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReleaseManager is a free data retrieval call binding the contract method 0x0d78573e.
//
// Solidity: function releaseManager() view returns(address)
func (_AppController *AppControllerSession) ReleaseManager() (common.Address, error) {
	return _AppController.Contract.ReleaseManager(&_AppController.CallOpts)
}

// ReleaseManager is a free data retrieval call binding the contract method 0x0d78573e.
//
// Solidity: function releaseManager() view returns(address)
func (_AppController *AppControllerCallerSession) ReleaseManager() (common.Address, error) {
	return _AppController.Contract.ReleaseManager(&_AppController.CallOpts)
}

// SafeTimelockFactory is a free data retrieval call binding the contract method 0xa03d2a81.
//
// Solidity: function safeTimelockFactory() view returns(address)
func (_AppController *AppControllerCaller) SafeTimelockFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "safeTimelockFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeTimelockFactory is a free data retrieval call binding the contract method 0xa03d2a81.
//
// Solidity: function safeTimelockFactory() view returns(address)
func (_AppController *AppControllerSession) SafeTimelockFactory() (common.Address, error) {
	return _AppController.Contract.SafeTimelockFactory(&_AppController.CallOpts)
}

// SafeTimelockFactory is a free data retrieval call binding the contract method 0xa03d2a81.
//
// Solidity: function safeTimelockFactory() view returns(address)
func (_AppController *AppControllerCallerSession) SafeTimelockFactory() (common.Address, error) {
	return _AppController.Contract.SafeTimelockFactory(&_AppController.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AppController *AppControllerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AppController *AppControllerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AppController.Contract.SupportsInterface(&_AppController.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AppController *AppControllerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AppController.Contract.SupportsInterface(&_AppController.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AppController *AppControllerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AppController.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AppController *AppControllerSession) Version() (string, error) {
	return _AppController.Contract.Version(&_AppController.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AppController *AppControllerCallerSession) Version() (string, error) {
	return _AppController.Contract.Version(&_AppController.CallOpts)
}

// CreateApp is a paid mutator transaction binding the contract method 0xa60daa8f.
//
// Solidity: function createApp(bytes32 salt, (((bytes32,string)[],uint32),bytes,bytes) release) returns(address app)
func (_AppController *AppControllerTransactor) CreateApp(opts *bind.TransactOpts, salt [32]byte, release IAppControllerRelease) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "createApp", salt, release)
}

// CreateApp is a paid mutator transaction binding the contract method 0xa60daa8f.
//
// Solidity: function createApp(bytes32 salt, (((bytes32,string)[],uint32),bytes,bytes) release) returns(address app)
func (_AppController *AppControllerSession) CreateApp(salt [32]byte, release IAppControllerRelease) (*types.Transaction, error) {
	return _AppController.Contract.CreateApp(&_AppController.TransactOpts, salt, release)
}

// CreateApp is a paid mutator transaction binding the contract method 0xa60daa8f.
//
// Solidity: function createApp(bytes32 salt, (((bytes32,string)[],uint32),bytes,bytes) release) returns(address app)
func (_AppController *AppControllerTransactorSession) CreateApp(salt [32]byte, release IAppControllerRelease) (*types.Transaction, error) {
	return _AppController.Contract.CreateApp(&_AppController.TransactOpts, salt, release)
}

// CreateAppForTeam is a paid mutator transaction binding the contract method 0x688c3ec3.
//
// Solidity: function createAppForTeam(address team, bytes32 salt, (((bytes32,string)[],uint32),bytes,bytes) release) returns(address app)
func (_AppController *AppControllerTransactor) CreateAppForTeam(opts *bind.TransactOpts, team common.Address, salt [32]byte, release IAppControllerRelease) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "createAppForTeam", team, salt, release)
}

// CreateAppForTeam is a paid mutator transaction binding the contract method 0x688c3ec3.
//
// Solidity: function createAppForTeam(address team, bytes32 salt, (((bytes32,string)[],uint32),bytes,bytes) release) returns(address app)
func (_AppController *AppControllerSession) CreateAppForTeam(team common.Address, salt [32]byte, release IAppControllerRelease) (*types.Transaction, error) {
	return _AppController.Contract.CreateAppForTeam(&_AppController.TransactOpts, team, salt, release)
}

// CreateAppForTeam is a paid mutator transaction binding the contract method 0x688c3ec3.
//
// Solidity: function createAppForTeam(address team, bytes32 salt, (((bytes32,string)[],uint32),bytes,bytes) release) returns(address app)
func (_AppController *AppControllerTransactorSession) CreateAppForTeam(team common.Address, salt [32]byte, release IAppControllerRelease) (*types.Transaction, error) {
	return _AppController.Contract.CreateAppForTeam(&_AppController.TransactOpts, team, salt, release)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AppController *AppControllerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AppController *AppControllerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.GrantRole(&_AppController.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AppController *AppControllerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.GrantRole(&_AppController.TransactOpts, role, account)
}

// GrantTeamRole is a paid mutator transaction binding the contract method 0x58f2c536.
//
// Solidity: function grantTeamRole(address team, uint8 role, address account) returns()
func (_AppController *AppControllerTransactor) GrantTeamRole(opts *bind.TransactOpts, team common.Address, role uint8, account common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "grantTeamRole", team, role, account)
}

// GrantTeamRole is a paid mutator transaction binding the contract method 0x58f2c536.
//
// Solidity: function grantTeamRole(address team, uint8 role, address account) returns()
func (_AppController *AppControllerSession) GrantTeamRole(team common.Address, role uint8, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.GrantTeamRole(&_AppController.TransactOpts, team, role, account)
}

// GrantTeamRole is a paid mutator transaction binding the contract method 0x58f2c536.
//
// Solidity: function grantTeamRole(address team, uint8 role, address account) returns()
func (_AppController *AppControllerTransactorSession) GrantTeamRole(team common.Address, role uint8, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.GrantTeamRole(&_AppController.TransactOpts, team, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_AppController *AppControllerTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_AppController *AppControllerSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _AppController.Contract.Initialize(&_AppController.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_AppController *AppControllerTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _AppController.Contract.Initialize(&_AppController.TransactOpts, admin)
}

// MigrateAdmins is a paid mutator transaction binding the contract method 0x00b73d4c.
//
// Solidity: function migrateAdmins(address[] apps) returns()
func (_AppController *AppControllerTransactor) MigrateAdmins(opts *bind.TransactOpts, apps []common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "migrateAdmins", apps)
}

// MigrateAdmins is a paid mutator transaction binding the contract method 0x00b73d4c.
//
// Solidity: function migrateAdmins(address[] apps) returns()
func (_AppController *AppControllerSession) MigrateAdmins(apps []common.Address) (*types.Transaction, error) {
	return _AppController.Contract.MigrateAdmins(&_AppController.TransactOpts, apps)
}

// MigrateAdmins is a paid mutator transaction binding the contract method 0x00b73d4c.
//
// Solidity: function migrateAdmins(address[] apps) returns()
func (_AppController *AppControllerTransactorSession) MigrateAdmins(apps []common.Address) (*types.Transaction, error) {
	return _AppController.Contract.MigrateAdmins(&_AppController.TransactOpts, apps)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AppController *AppControllerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AppController *AppControllerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.RenounceRole(&_AppController.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AppController *AppControllerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.RenounceRole(&_AppController.TransactOpts, role, account)
}

// RenounceTeamRole is a paid mutator transaction binding the contract method 0x17f4c90b.
//
// Solidity: function renounceTeamRole(address team, uint8 role) returns()
func (_AppController *AppControllerTransactor) RenounceTeamRole(opts *bind.TransactOpts, team common.Address, role uint8) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "renounceTeamRole", team, role)
}

// RenounceTeamRole is a paid mutator transaction binding the contract method 0x17f4c90b.
//
// Solidity: function renounceTeamRole(address team, uint8 role) returns()
func (_AppController *AppControllerSession) RenounceTeamRole(team common.Address, role uint8) (*types.Transaction, error) {
	return _AppController.Contract.RenounceTeamRole(&_AppController.TransactOpts, team, role)
}

// RenounceTeamRole is a paid mutator transaction binding the contract method 0x17f4c90b.
//
// Solidity: function renounceTeamRole(address team, uint8 role) returns()
func (_AppController *AppControllerTransactorSession) RenounceTeamRole(team common.Address, role uint8) (*types.Transaction, error) {
	return _AppController.Contract.RenounceTeamRole(&_AppController.TransactOpts, team, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AppController *AppControllerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AppController *AppControllerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.RevokeRole(&_AppController.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AppController *AppControllerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.RevokeRole(&_AppController.TransactOpts, role, account)
}

// RevokeTeamRole is a paid mutator transaction binding the contract method 0x11b652e2.
//
// Solidity: function revokeTeamRole(address team, uint8 role, address account) returns()
func (_AppController *AppControllerTransactor) RevokeTeamRole(opts *bind.TransactOpts, team common.Address, role uint8, account common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "revokeTeamRole", team, role, account)
}

// RevokeTeamRole is a paid mutator transaction binding the contract method 0x11b652e2.
//
// Solidity: function revokeTeamRole(address team, uint8 role, address account) returns()
func (_AppController *AppControllerSession) RevokeTeamRole(team common.Address, role uint8, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.RevokeTeamRole(&_AppController.TransactOpts, team, role, account)
}

// RevokeTeamRole is a paid mutator transaction binding the contract method 0x11b652e2.
//
// Solidity: function revokeTeamRole(address team, uint8 role, address account) returns()
func (_AppController *AppControllerTransactorSession) RevokeTeamRole(team common.Address, role uint8, account common.Address) (*types.Transaction, error) {
	return _AppController.Contract.RevokeTeamRole(&_AppController.TransactOpts, team, role, account)
}

// SetMaxActiveAppsPerUser is a paid mutator transaction binding the contract method 0xd49fec2b.
//
// Solidity: function setMaxActiveAppsPerUser(address user, uint32 limit) returns()
func (_AppController *AppControllerTransactor) SetMaxActiveAppsPerUser(opts *bind.TransactOpts, user common.Address, limit uint32) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "setMaxActiveAppsPerUser", user, limit)
}

// SetMaxActiveAppsPerUser is a paid mutator transaction binding the contract method 0xd49fec2b.
//
// Solidity: function setMaxActiveAppsPerUser(address user, uint32 limit) returns()
func (_AppController *AppControllerSession) SetMaxActiveAppsPerUser(user common.Address, limit uint32) (*types.Transaction, error) {
	return _AppController.Contract.SetMaxActiveAppsPerUser(&_AppController.TransactOpts, user, limit)
}

// SetMaxActiveAppsPerUser is a paid mutator transaction binding the contract method 0xd49fec2b.
//
// Solidity: function setMaxActiveAppsPerUser(address user, uint32 limit) returns()
func (_AppController *AppControllerTransactorSession) SetMaxActiveAppsPerUser(user common.Address, limit uint32) (*types.Transaction, error) {
	return _AppController.Contract.SetMaxActiveAppsPerUser(&_AppController.TransactOpts, user, limit)
}

// SetMaxGlobalActiveApps is a paid mutator transaction binding the contract method 0xb438c141.
//
// Solidity: function setMaxGlobalActiveApps(uint32 limit) returns()
func (_AppController *AppControllerTransactor) SetMaxGlobalActiveApps(opts *bind.TransactOpts, limit uint32) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "setMaxGlobalActiveApps", limit)
}

// SetMaxGlobalActiveApps is a paid mutator transaction binding the contract method 0xb438c141.
//
// Solidity: function setMaxGlobalActiveApps(uint32 limit) returns()
func (_AppController *AppControllerSession) SetMaxGlobalActiveApps(limit uint32) (*types.Transaction, error) {
	return _AppController.Contract.SetMaxGlobalActiveApps(&_AppController.TransactOpts, limit)
}

// SetMaxGlobalActiveApps is a paid mutator transaction binding the contract method 0xb438c141.
//
// Solidity: function setMaxGlobalActiveApps(uint32 limit) returns()
func (_AppController *AppControllerTransactorSession) SetMaxGlobalActiveApps(limit uint32) (*types.Transaction, error) {
	return _AppController.Contract.SetMaxGlobalActiveApps(&_AppController.TransactOpts, limit)
}

// StartApp is a paid mutator transaction binding the contract method 0x4dd8ef81.
//
// Solidity: function startApp(address app) returns()
func (_AppController *AppControllerTransactor) StartApp(opts *bind.TransactOpts, app common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "startApp", app)
}

// StartApp is a paid mutator transaction binding the contract method 0x4dd8ef81.
//
// Solidity: function startApp(address app) returns()
func (_AppController *AppControllerSession) StartApp(app common.Address) (*types.Transaction, error) {
	return _AppController.Contract.StartApp(&_AppController.TransactOpts, app)
}

// StartApp is a paid mutator transaction binding the contract method 0x4dd8ef81.
//
// Solidity: function startApp(address app) returns()
func (_AppController *AppControllerTransactorSession) StartApp(app common.Address) (*types.Transaction, error) {
	return _AppController.Contract.StartApp(&_AppController.TransactOpts, app)
}

// StopApp is a paid mutator transaction binding the contract method 0xe70b81b1.
//
// Solidity: function stopApp(address app) returns()
func (_AppController *AppControllerTransactor) StopApp(opts *bind.TransactOpts, app common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "stopApp", app)
}

// StopApp is a paid mutator transaction binding the contract method 0xe70b81b1.
//
// Solidity: function stopApp(address app) returns()
func (_AppController *AppControllerSession) StopApp(app common.Address) (*types.Transaction, error) {
	return _AppController.Contract.StopApp(&_AppController.TransactOpts, app)
}

// StopApp is a paid mutator transaction binding the contract method 0xe70b81b1.
//
// Solidity: function stopApp(address app) returns()
func (_AppController *AppControllerTransactorSession) StopApp(app common.Address) (*types.Transaction, error) {
	return _AppController.Contract.StopApp(&_AppController.TransactOpts, app)
}

// Suspend is a paid mutator transaction binding the contract method 0xcb1e6ff7.
//
// Solidity: function suspend(address account, address[] apps) returns()
func (_AppController *AppControllerTransactor) Suspend(opts *bind.TransactOpts, account common.Address, apps []common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "suspend", account, apps)
}

// Suspend is a paid mutator transaction binding the contract method 0xcb1e6ff7.
//
// Solidity: function suspend(address account, address[] apps) returns()
func (_AppController *AppControllerSession) Suspend(account common.Address, apps []common.Address) (*types.Transaction, error) {
	return _AppController.Contract.Suspend(&_AppController.TransactOpts, account, apps)
}

// Suspend is a paid mutator transaction binding the contract method 0xcb1e6ff7.
//
// Solidity: function suspend(address account, address[] apps) returns()
func (_AppController *AppControllerTransactorSession) Suspend(account common.Address, apps []common.Address) (*types.Transaction, error) {
	return _AppController.Contract.Suspend(&_AppController.TransactOpts, account, apps)
}

// TerminateApp is a paid mutator transaction binding the contract method 0x90e2196b.
//
// Solidity: function terminateApp(address app) returns()
func (_AppController *AppControllerTransactor) TerminateApp(opts *bind.TransactOpts, app common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "terminateApp", app)
}

// TerminateApp is a paid mutator transaction binding the contract method 0x90e2196b.
//
// Solidity: function terminateApp(address app) returns()
func (_AppController *AppControllerSession) TerminateApp(app common.Address) (*types.Transaction, error) {
	return _AppController.Contract.TerminateApp(&_AppController.TransactOpts, app)
}

// TerminateApp is a paid mutator transaction binding the contract method 0x90e2196b.
//
// Solidity: function terminateApp(address app) returns()
func (_AppController *AppControllerTransactorSession) TerminateApp(app common.Address) (*types.Transaction, error) {
	return _AppController.Contract.TerminateApp(&_AppController.TransactOpts, app)
}

// TerminateAppByAdmin is a paid mutator transaction binding the contract method 0xd66f7771.
//
// Solidity: function terminateAppByAdmin(address app) returns()
func (_AppController *AppControllerTransactor) TerminateAppByAdmin(opts *bind.TransactOpts, app common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "terminateAppByAdmin", app)
}

// TerminateAppByAdmin is a paid mutator transaction binding the contract method 0xd66f7771.
//
// Solidity: function terminateAppByAdmin(address app) returns()
func (_AppController *AppControllerSession) TerminateAppByAdmin(app common.Address) (*types.Transaction, error) {
	return _AppController.Contract.TerminateAppByAdmin(&_AppController.TransactOpts, app)
}

// TerminateAppByAdmin is a paid mutator transaction binding the contract method 0xd66f7771.
//
// Solidity: function terminateAppByAdmin(address app) returns()
func (_AppController *AppControllerTransactorSession) TerminateAppByAdmin(app common.Address) (*types.Transaction, error) {
	return _AppController.Contract.TerminateAppByAdmin(&_AppController.TransactOpts, app)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address app, address newOwner) returns()
func (_AppController *AppControllerTransactor) TransferOwnership(opts *bind.TransactOpts, app common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "transferOwnership", app, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address app, address newOwner) returns()
func (_AppController *AppControllerSession) TransferOwnership(app common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _AppController.Contract.TransferOwnership(&_AppController.TransactOpts, app, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address app, address newOwner) returns()
func (_AppController *AppControllerTransactorSession) TransferOwnership(app common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _AppController.Contract.TransferOwnership(&_AppController.TransactOpts, app, newOwner)
}

// UpdateAppMetadataURI is a paid mutator transaction binding the contract method 0x65aa9a65.
//
// Solidity: function updateAppMetadataURI(address app, string metadataURI) returns()
func (_AppController *AppControllerTransactor) UpdateAppMetadataURI(opts *bind.TransactOpts, app common.Address, metadataURI string) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "updateAppMetadataURI", app, metadataURI)
}

// UpdateAppMetadataURI is a paid mutator transaction binding the contract method 0x65aa9a65.
//
// Solidity: function updateAppMetadataURI(address app, string metadataURI) returns()
func (_AppController *AppControllerSession) UpdateAppMetadataURI(app common.Address, metadataURI string) (*types.Transaction, error) {
	return _AppController.Contract.UpdateAppMetadataURI(&_AppController.TransactOpts, app, metadataURI)
}

// UpdateAppMetadataURI is a paid mutator transaction binding the contract method 0x65aa9a65.
//
// Solidity: function updateAppMetadataURI(address app, string metadataURI) returns()
func (_AppController *AppControllerTransactorSession) UpdateAppMetadataURI(app common.Address, metadataURI string) (*types.Transaction, error) {
	return _AppController.Contract.UpdateAppMetadataURI(&_AppController.TransactOpts, app, metadataURI)
}

// UpgradeApp is a paid mutator transaction binding the contract method 0xd80a956b.
//
// Solidity: function upgradeApp(address app, (((bytes32,string)[],uint32),bytes,bytes) release) returns(uint256)
func (_AppController *AppControllerTransactor) UpgradeApp(opts *bind.TransactOpts, app common.Address, release IAppControllerRelease) (*types.Transaction, error) {
	return _AppController.contract.Transact(opts, "upgradeApp", app, release)
}

// UpgradeApp is a paid mutator transaction binding the contract method 0xd80a956b.
//
// Solidity: function upgradeApp(address app, (((bytes32,string)[],uint32),bytes,bytes) release) returns(uint256)
func (_AppController *AppControllerSession) UpgradeApp(app common.Address, release IAppControllerRelease) (*types.Transaction, error) {
	return _AppController.Contract.UpgradeApp(&_AppController.TransactOpts, app, release)
}

// UpgradeApp is a paid mutator transaction binding the contract method 0xd80a956b.
//
// Solidity: function upgradeApp(address app, (((bytes32,string)[],uint32),bytes,bytes) release) returns(uint256)
func (_AppController *AppControllerTransactorSession) UpgradeApp(app common.Address, release IAppControllerRelease) (*types.Transaction, error) {
	return _AppController.Contract.UpgradeApp(&_AppController.TransactOpts, app, release)
}

// AppControllerAppCreatedIterator is returned from FilterAppCreated and is used to iterate over the raw logs and unpacked data for AppCreated events raised by the AppController contract.
type AppControllerAppCreatedIterator struct {
	Event *AppControllerAppCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerAppCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerAppCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerAppCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerAppCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerAppCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerAppCreated represents a AppCreated event raised by the AppController contract.
type AppControllerAppCreated struct {
	Owner         common.Address
	App           common.Address
	OperatorSetId uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAppCreated is a free log retrieval operation binding the contract event 0xe9e6e5409b80e2dab7d38194fb370b2e8045a1af28177b94ddcf19d2495d7589.
//
// Solidity: event AppCreated(address indexed owner, address indexed app, uint32 operatorSetId)
func (_AppController *AppControllerFilterer) FilterAppCreated(opts *bind.FilterOpts, owner []common.Address, app []common.Address) (*AppControllerAppCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "AppCreated", ownerRule, appRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerAppCreatedIterator{contract: _AppController.contract, event: "AppCreated", logs: logs, sub: sub}, nil
}

// WatchAppCreated is a free log subscription operation binding the contract event 0xe9e6e5409b80e2dab7d38194fb370b2e8045a1af28177b94ddcf19d2495d7589.
//
// Solidity: event AppCreated(address indexed owner, address indexed app, uint32 operatorSetId)
func (_AppController *AppControllerFilterer) WatchAppCreated(opts *bind.WatchOpts, sink chan<- *AppControllerAppCreated, owner []common.Address, app []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "AppCreated", ownerRule, appRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerAppCreated)
				if err := _AppController.contract.UnpackLog(event, "AppCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppCreated is a log parse operation binding the contract event 0xe9e6e5409b80e2dab7d38194fb370b2e8045a1af28177b94ddcf19d2495d7589.
//
// Solidity: event AppCreated(address indexed owner, address indexed app, uint32 operatorSetId)
func (_AppController *AppControllerFilterer) ParseAppCreated(log types.Log) (*AppControllerAppCreated, error) {
	event := new(AppControllerAppCreated)
	if err := _AppController.contract.UnpackLog(event, "AppCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerAppMetadataURIUpdatedIterator is returned from FilterAppMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AppMetadataURIUpdated events raised by the AppController contract.
type AppControllerAppMetadataURIUpdatedIterator struct {
	Event *AppControllerAppMetadataURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerAppMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerAppMetadataURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerAppMetadataURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerAppMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerAppMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerAppMetadataURIUpdated represents a AppMetadataURIUpdated event raised by the AppController contract.
type AppControllerAppMetadataURIUpdated struct {
	App         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAppMetadataURIUpdated is a free log retrieval operation binding the contract event 0xc874c4e2852c19bd03d42a5232ce93e85f2da325dddac46c2472043b7f0d7e8d.
//
// Solidity: event AppMetadataURIUpdated(address indexed app, string metadataURI)
func (_AppController *AppControllerFilterer) FilterAppMetadataURIUpdated(opts *bind.FilterOpts, app []common.Address) (*AppControllerAppMetadataURIUpdatedIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "AppMetadataURIUpdated", appRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerAppMetadataURIUpdatedIterator{contract: _AppController.contract, event: "AppMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAppMetadataURIUpdated is a free log subscription operation binding the contract event 0xc874c4e2852c19bd03d42a5232ce93e85f2da325dddac46c2472043b7f0d7e8d.
//
// Solidity: event AppMetadataURIUpdated(address indexed app, string metadataURI)
func (_AppController *AppControllerFilterer) WatchAppMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *AppControllerAppMetadataURIUpdated, app []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "AppMetadataURIUpdated", appRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerAppMetadataURIUpdated)
				if err := _AppController.contract.UnpackLog(event, "AppMetadataURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppMetadataURIUpdated is a log parse operation binding the contract event 0xc874c4e2852c19bd03d42a5232ce93e85f2da325dddac46c2472043b7f0d7e8d.
//
// Solidity: event AppMetadataURIUpdated(address indexed app, string metadataURI)
func (_AppController *AppControllerFilterer) ParseAppMetadataURIUpdated(log types.Log) (*AppControllerAppMetadataURIUpdated, error) {
	event := new(AppControllerAppMetadataURIUpdated)
	if err := _AppController.contract.UnpackLog(event, "AppMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerAppOwnershipTransferredIterator is returned from FilterAppOwnershipTransferred and is used to iterate over the raw logs and unpacked data for AppOwnershipTransferred events raised by the AppController contract.
type AppControllerAppOwnershipTransferredIterator struct {
	Event *AppControllerAppOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerAppOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerAppOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerAppOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerAppOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerAppOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerAppOwnershipTransferred represents a AppOwnershipTransferred event raised by the AppController contract.
type AppControllerAppOwnershipTransferred struct {
	App           common.Address
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAppOwnershipTransferred is a free log retrieval operation binding the contract event 0x3fa03516e5ee455b2d2779f21b254735e2c1f82cf338619c1b96816df2a467a4.
//
// Solidity: event AppOwnershipTransferred(address indexed app, address indexed previousOwner, address indexed newOwner)
func (_AppController *AppControllerFilterer) FilterAppOwnershipTransferred(opts *bind.FilterOpts, app []common.Address, previousOwner []common.Address, newOwner []common.Address) (*AppControllerAppOwnershipTransferredIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "AppOwnershipTransferred", appRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerAppOwnershipTransferredIterator{contract: _AppController.contract, event: "AppOwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchAppOwnershipTransferred is a free log subscription operation binding the contract event 0x3fa03516e5ee455b2d2779f21b254735e2c1f82cf338619c1b96816df2a467a4.
//
// Solidity: event AppOwnershipTransferred(address indexed app, address indexed previousOwner, address indexed newOwner)
func (_AppController *AppControllerFilterer) WatchAppOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AppControllerAppOwnershipTransferred, app []common.Address, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "AppOwnershipTransferred", appRule, previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerAppOwnershipTransferred)
				if err := _AppController.contract.UnpackLog(event, "AppOwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppOwnershipTransferred is a log parse operation binding the contract event 0x3fa03516e5ee455b2d2779f21b254735e2c1f82cf338619c1b96816df2a467a4.
//
// Solidity: event AppOwnershipTransferred(address indexed app, address indexed previousOwner, address indexed newOwner)
func (_AppController *AppControllerFilterer) ParseAppOwnershipTransferred(log types.Log) (*AppControllerAppOwnershipTransferred, error) {
	event := new(AppControllerAppOwnershipTransferred)
	if err := _AppController.contract.UnpackLog(event, "AppOwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerAppStartedIterator is returned from FilterAppStarted and is used to iterate over the raw logs and unpacked data for AppStarted events raised by the AppController contract.
type AppControllerAppStartedIterator struct {
	Event *AppControllerAppStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerAppStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerAppStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerAppStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerAppStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerAppStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerAppStarted represents a AppStarted event raised by the AppController contract.
type AppControllerAppStarted struct {
	App common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppStarted is a free log retrieval operation binding the contract event 0x8426864999f5f9244a203f75107b05f8b474838bc43ff0d8abe18ee5f4427513.
//
// Solidity: event AppStarted(address indexed app)
func (_AppController *AppControllerFilterer) FilterAppStarted(opts *bind.FilterOpts, app []common.Address) (*AppControllerAppStartedIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "AppStarted", appRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerAppStartedIterator{contract: _AppController.contract, event: "AppStarted", logs: logs, sub: sub}, nil
}

// WatchAppStarted is a free log subscription operation binding the contract event 0x8426864999f5f9244a203f75107b05f8b474838bc43ff0d8abe18ee5f4427513.
//
// Solidity: event AppStarted(address indexed app)
func (_AppController *AppControllerFilterer) WatchAppStarted(opts *bind.WatchOpts, sink chan<- *AppControllerAppStarted, app []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "AppStarted", appRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerAppStarted)
				if err := _AppController.contract.UnpackLog(event, "AppStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppStarted is a log parse operation binding the contract event 0x8426864999f5f9244a203f75107b05f8b474838bc43ff0d8abe18ee5f4427513.
//
// Solidity: event AppStarted(address indexed app)
func (_AppController *AppControllerFilterer) ParseAppStarted(log types.Log) (*AppControllerAppStarted, error) {
	event := new(AppControllerAppStarted)
	if err := _AppController.contract.UnpackLog(event, "AppStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerAppStoppedIterator is returned from FilterAppStopped and is used to iterate over the raw logs and unpacked data for AppStopped events raised by the AppController contract.
type AppControllerAppStoppedIterator struct {
	Event *AppControllerAppStopped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerAppStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerAppStopped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerAppStopped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerAppStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerAppStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerAppStopped represents a AppStopped event raised by the AppController contract.
type AppControllerAppStopped struct {
	App common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppStopped is a free log retrieval operation binding the contract event 0x143b9e3b5aa0fb8cf19b1bdb2eb8d7a94657d4b505070c667433d322b90b1f78.
//
// Solidity: event AppStopped(address indexed app)
func (_AppController *AppControllerFilterer) FilterAppStopped(opts *bind.FilterOpts, app []common.Address) (*AppControllerAppStoppedIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "AppStopped", appRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerAppStoppedIterator{contract: _AppController.contract, event: "AppStopped", logs: logs, sub: sub}, nil
}

// WatchAppStopped is a free log subscription operation binding the contract event 0x143b9e3b5aa0fb8cf19b1bdb2eb8d7a94657d4b505070c667433d322b90b1f78.
//
// Solidity: event AppStopped(address indexed app)
func (_AppController *AppControllerFilterer) WatchAppStopped(opts *bind.WatchOpts, sink chan<- *AppControllerAppStopped, app []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "AppStopped", appRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerAppStopped)
				if err := _AppController.contract.UnpackLog(event, "AppStopped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppStopped is a log parse operation binding the contract event 0x143b9e3b5aa0fb8cf19b1bdb2eb8d7a94657d4b505070c667433d322b90b1f78.
//
// Solidity: event AppStopped(address indexed app)
func (_AppController *AppControllerFilterer) ParseAppStopped(log types.Log) (*AppControllerAppStopped, error) {
	event := new(AppControllerAppStopped)
	if err := _AppController.contract.UnpackLog(event, "AppStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerAppSuspendedIterator is returned from FilterAppSuspended and is used to iterate over the raw logs and unpacked data for AppSuspended events raised by the AppController contract.
type AppControllerAppSuspendedIterator struct {
	Event *AppControllerAppSuspended // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerAppSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerAppSuspended)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerAppSuspended)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerAppSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerAppSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerAppSuspended represents a AppSuspended event raised by the AppController contract.
type AppControllerAppSuspended struct {
	App common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppSuspended is a free log retrieval operation binding the contract event 0xbb85757a09bd73484c9fb6857c3e9f8dbf92c3e22ce1655213c9481dacf81c62.
//
// Solidity: event AppSuspended(address indexed app)
func (_AppController *AppControllerFilterer) FilterAppSuspended(opts *bind.FilterOpts, app []common.Address) (*AppControllerAppSuspendedIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "AppSuspended", appRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerAppSuspendedIterator{contract: _AppController.contract, event: "AppSuspended", logs: logs, sub: sub}, nil
}

// WatchAppSuspended is a free log subscription operation binding the contract event 0xbb85757a09bd73484c9fb6857c3e9f8dbf92c3e22ce1655213c9481dacf81c62.
//
// Solidity: event AppSuspended(address indexed app)
func (_AppController *AppControllerFilterer) WatchAppSuspended(opts *bind.WatchOpts, sink chan<- *AppControllerAppSuspended, app []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "AppSuspended", appRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerAppSuspended)
				if err := _AppController.contract.UnpackLog(event, "AppSuspended", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppSuspended is a log parse operation binding the contract event 0xbb85757a09bd73484c9fb6857c3e9f8dbf92c3e22ce1655213c9481dacf81c62.
//
// Solidity: event AppSuspended(address indexed app)
func (_AppController *AppControllerFilterer) ParseAppSuspended(log types.Log) (*AppControllerAppSuspended, error) {
	event := new(AppControllerAppSuspended)
	if err := _AppController.contract.UnpackLog(event, "AppSuspended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerAppTerminatedIterator is returned from FilterAppTerminated and is used to iterate over the raw logs and unpacked data for AppTerminated events raised by the AppController contract.
type AppControllerAppTerminatedIterator struct {
	Event *AppControllerAppTerminated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerAppTerminatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerAppTerminated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerAppTerminated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerAppTerminatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerAppTerminatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerAppTerminated represents a AppTerminated event raised by the AppController contract.
type AppControllerAppTerminated struct {
	App common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppTerminated is a free log retrieval operation binding the contract event 0x2465e924a1e9a02843493e65702fd09152c8a631caf7b9b58e4b01ee20ced416.
//
// Solidity: event AppTerminated(address indexed app)
func (_AppController *AppControllerFilterer) FilterAppTerminated(opts *bind.FilterOpts, app []common.Address) (*AppControllerAppTerminatedIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "AppTerminated", appRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerAppTerminatedIterator{contract: _AppController.contract, event: "AppTerminated", logs: logs, sub: sub}, nil
}

// WatchAppTerminated is a free log subscription operation binding the contract event 0x2465e924a1e9a02843493e65702fd09152c8a631caf7b9b58e4b01ee20ced416.
//
// Solidity: event AppTerminated(address indexed app)
func (_AppController *AppControllerFilterer) WatchAppTerminated(opts *bind.WatchOpts, sink chan<- *AppControllerAppTerminated, app []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "AppTerminated", appRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerAppTerminated)
				if err := _AppController.contract.UnpackLog(event, "AppTerminated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppTerminated is a log parse operation binding the contract event 0x2465e924a1e9a02843493e65702fd09152c8a631caf7b9b58e4b01ee20ced416.
//
// Solidity: event AppTerminated(address indexed app)
func (_AppController *AppControllerFilterer) ParseAppTerminated(log types.Log) (*AppControllerAppTerminated, error) {
	event := new(AppControllerAppTerminated)
	if err := _AppController.contract.UnpackLog(event, "AppTerminated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerAppTerminatedByAdminIterator is returned from FilterAppTerminatedByAdmin and is used to iterate over the raw logs and unpacked data for AppTerminatedByAdmin events raised by the AppController contract.
type AppControllerAppTerminatedByAdminIterator struct {
	Event *AppControllerAppTerminatedByAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerAppTerminatedByAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerAppTerminatedByAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerAppTerminatedByAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerAppTerminatedByAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerAppTerminatedByAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerAppTerminatedByAdmin represents a AppTerminatedByAdmin event raised by the AppController contract.
type AppControllerAppTerminatedByAdmin struct {
	App common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppTerminatedByAdmin is a free log retrieval operation binding the contract event 0x4636ad9a0a072d9c8b05a5bcda7d33256c5401515e612c25c13273b378cb2809.
//
// Solidity: event AppTerminatedByAdmin(address indexed app)
func (_AppController *AppControllerFilterer) FilterAppTerminatedByAdmin(opts *bind.FilterOpts, app []common.Address) (*AppControllerAppTerminatedByAdminIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "AppTerminatedByAdmin", appRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerAppTerminatedByAdminIterator{contract: _AppController.contract, event: "AppTerminatedByAdmin", logs: logs, sub: sub}, nil
}

// WatchAppTerminatedByAdmin is a free log subscription operation binding the contract event 0x4636ad9a0a072d9c8b05a5bcda7d33256c5401515e612c25c13273b378cb2809.
//
// Solidity: event AppTerminatedByAdmin(address indexed app)
func (_AppController *AppControllerFilterer) WatchAppTerminatedByAdmin(opts *bind.WatchOpts, sink chan<- *AppControllerAppTerminatedByAdmin, app []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "AppTerminatedByAdmin", appRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerAppTerminatedByAdmin)
				if err := _AppController.contract.UnpackLog(event, "AppTerminatedByAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppTerminatedByAdmin is a log parse operation binding the contract event 0x4636ad9a0a072d9c8b05a5bcda7d33256c5401515e612c25c13273b378cb2809.
//
// Solidity: event AppTerminatedByAdmin(address indexed app)
func (_AppController *AppControllerFilterer) ParseAppTerminatedByAdmin(log types.Log) (*AppControllerAppTerminatedByAdmin, error) {
	event := new(AppControllerAppTerminatedByAdmin)
	if err := _AppController.contract.UnpackLog(event, "AppTerminatedByAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerAppUpgradedIterator is returned from FilterAppUpgraded and is used to iterate over the raw logs and unpacked data for AppUpgraded events raised by the AppController contract.
type AppControllerAppUpgradedIterator struct {
	Event *AppControllerAppUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerAppUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerAppUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerAppUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerAppUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerAppUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerAppUpgraded represents a AppUpgraded event raised by the AppController contract.
type AppControllerAppUpgraded struct {
	App          common.Address
	RmsReleaseId *big.Int
	Release      IAppControllerRelease
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAppUpgraded is a free log retrieval operation binding the contract event 0xb5d2b7df0352a87a587a430bfceab43859b3b3c404f0425c4c1ab0675ed57b94.
//
// Solidity: event AppUpgraded(address indexed app, uint256 rmsReleaseId, (((bytes32,string)[],uint32),bytes,bytes) release)
func (_AppController *AppControllerFilterer) FilterAppUpgraded(opts *bind.FilterOpts, app []common.Address) (*AppControllerAppUpgradedIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "AppUpgraded", appRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerAppUpgradedIterator{contract: _AppController.contract, event: "AppUpgraded", logs: logs, sub: sub}, nil
}

// WatchAppUpgraded is a free log subscription operation binding the contract event 0xb5d2b7df0352a87a587a430bfceab43859b3b3c404f0425c4c1ab0675ed57b94.
//
// Solidity: event AppUpgraded(address indexed app, uint256 rmsReleaseId, (((bytes32,string)[],uint32),bytes,bytes) release)
func (_AppController *AppControllerFilterer) WatchAppUpgraded(opts *bind.WatchOpts, sink chan<- *AppControllerAppUpgraded, app []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "AppUpgraded", appRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerAppUpgraded)
				if err := _AppController.contract.UnpackLog(event, "AppUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppUpgraded is a log parse operation binding the contract event 0xb5d2b7df0352a87a587a430bfceab43859b3b3c404f0425c4c1ab0675ed57b94.
//
// Solidity: event AppUpgraded(address indexed app, uint256 rmsReleaseId, (((bytes32,string)[],uint32),bytes,bytes) release)
func (_AppController *AppControllerFilterer) ParseAppUpgraded(log types.Log) (*AppControllerAppUpgraded, error) {
	event := new(AppControllerAppUpgraded)
	if err := _AppController.contract.UnpackLog(event, "AppUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerGlobalMaxActiveAppsSetIterator is returned from FilterGlobalMaxActiveAppsSet and is used to iterate over the raw logs and unpacked data for GlobalMaxActiveAppsSet events raised by the AppController contract.
type AppControllerGlobalMaxActiveAppsSetIterator struct {
	Event *AppControllerGlobalMaxActiveAppsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerGlobalMaxActiveAppsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerGlobalMaxActiveAppsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerGlobalMaxActiveAppsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerGlobalMaxActiveAppsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerGlobalMaxActiveAppsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerGlobalMaxActiveAppsSet represents a GlobalMaxActiveAppsSet event raised by the AppController contract.
type AppControllerGlobalMaxActiveAppsSet struct {
	Limit uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGlobalMaxActiveAppsSet is a free log retrieval operation binding the contract event 0xc3c99803ad2e80be7382791ca0b23147a94739b9e66a656b9c8c47f9d6963414.
//
// Solidity: event GlobalMaxActiveAppsSet(uint32 limit)
func (_AppController *AppControllerFilterer) FilterGlobalMaxActiveAppsSet(opts *bind.FilterOpts) (*AppControllerGlobalMaxActiveAppsSetIterator, error) {

	logs, sub, err := _AppController.contract.FilterLogs(opts, "GlobalMaxActiveAppsSet")
	if err != nil {
		return nil, err
	}
	return &AppControllerGlobalMaxActiveAppsSetIterator{contract: _AppController.contract, event: "GlobalMaxActiveAppsSet", logs: logs, sub: sub}, nil
}

// WatchGlobalMaxActiveAppsSet is a free log subscription operation binding the contract event 0xc3c99803ad2e80be7382791ca0b23147a94739b9e66a656b9c8c47f9d6963414.
//
// Solidity: event GlobalMaxActiveAppsSet(uint32 limit)
func (_AppController *AppControllerFilterer) WatchGlobalMaxActiveAppsSet(opts *bind.WatchOpts, sink chan<- *AppControllerGlobalMaxActiveAppsSet) (event.Subscription, error) {

	logs, sub, err := _AppController.contract.WatchLogs(opts, "GlobalMaxActiveAppsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerGlobalMaxActiveAppsSet)
				if err := _AppController.contract.UnpackLog(event, "GlobalMaxActiveAppsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGlobalMaxActiveAppsSet is a log parse operation binding the contract event 0xc3c99803ad2e80be7382791ca0b23147a94739b9e66a656b9c8c47f9d6963414.
//
// Solidity: event GlobalMaxActiveAppsSet(uint32 limit)
func (_AppController *AppControllerFilterer) ParseGlobalMaxActiveAppsSet(log types.Log) (*AppControllerGlobalMaxActiveAppsSet, error) {
	event := new(AppControllerGlobalMaxActiveAppsSet)
	if err := _AppController.contract.UnpackLog(event, "GlobalMaxActiveAppsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AppController contract.
type AppControllerInitializedIterator struct {
	Event *AppControllerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerInitialized represents a Initialized event raised by the AppController contract.
type AppControllerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AppController *AppControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AppControllerInitializedIterator, error) {

	logs, sub, err := _AppController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AppControllerInitializedIterator{contract: _AppController.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AppController *AppControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AppControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _AppController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerInitialized)
				if err := _AppController.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AppController *AppControllerFilterer) ParseInitialized(log types.Log) (*AppControllerInitialized, error) {
	event := new(AppControllerInitialized)
	if err := _AppController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerMaxActiveAppsSetIterator is returned from FilterMaxActiveAppsSet and is used to iterate over the raw logs and unpacked data for MaxActiveAppsSet events raised by the AppController contract.
type AppControllerMaxActiveAppsSetIterator struct {
	Event *AppControllerMaxActiveAppsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerMaxActiveAppsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerMaxActiveAppsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerMaxActiveAppsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerMaxActiveAppsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerMaxActiveAppsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerMaxActiveAppsSet represents a MaxActiveAppsSet event raised by the AppController contract.
type AppControllerMaxActiveAppsSet struct {
	User  common.Address
	Limit uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMaxActiveAppsSet is a free log retrieval operation binding the contract event 0x708c955faa87ece750a9700c153a20797167e2fcbde9fca9186801f874663a99.
//
// Solidity: event MaxActiveAppsSet(address indexed user, uint32 limit)
func (_AppController *AppControllerFilterer) FilterMaxActiveAppsSet(opts *bind.FilterOpts, user []common.Address) (*AppControllerMaxActiveAppsSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "MaxActiveAppsSet", userRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerMaxActiveAppsSetIterator{contract: _AppController.contract, event: "MaxActiveAppsSet", logs: logs, sub: sub}, nil
}

// WatchMaxActiveAppsSet is a free log subscription operation binding the contract event 0x708c955faa87ece750a9700c153a20797167e2fcbde9fca9186801f874663a99.
//
// Solidity: event MaxActiveAppsSet(address indexed user, uint32 limit)
func (_AppController *AppControllerFilterer) WatchMaxActiveAppsSet(opts *bind.WatchOpts, sink chan<- *AppControllerMaxActiveAppsSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "MaxActiveAppsSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerMaxActiveAppsSet)
				if err := _AppController.contract.UnpackLog(event, "MaxActiveAppsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxActiveAppsSet is a log parse operation binding the contract event 0x708c955faa87ece750a9700c153a20797167e2fcbde9fca9186801f874663a99.
//
// Solidity: event MaxActiveAppsSet(address indexed user, uint32 limit)
func (_AppController *AppControllerFilterer) ParseMaxActiveAppsSet(log types.Log) (*AppControllerMaxActiveAppsSet, error) {
	event := new(AppControllerMaxActiveAppsSet)
	if err := _AppController.contract.UnpackLog(event, "MaxActiveAppsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AppController contract.
type AppControllerRoleAdminChangedIterator struct {
	Event *AppControllerRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerRoleAdminChanged represents a RoleAdminChanged event raised by the AppController contract.
type AppControllerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AppController *AppControllerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AppControllerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerRoleAdminChangedIterator{contract: _AppController.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AppController *AppControllerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AppControllerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerRoleAdminChanged)
				if err := _AppController.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AppController *AppControllerFilterer) ParseRoleAdminChanged(log types.Log) (*AppControllerRoleAdminChanged, error) {
	event := new(AppControllerRoleAdminChanged)
	if err := _AppController.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AppController contract.
type AppControllerRoleGrantedIterator struct {
	Event *AppControllerRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerRoleGranted represents a RoleGranted event raised by the AppController contract.
type AppControllerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppController *AppControllerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AppControllerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerRoleGrantedIterator{contract: _AppController.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppController *AppControllerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AppControllerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerRoleGranted)
				if err := _AppController.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppController *AppControllerFilterer) ParseRoleGranted(log types.Log) (*AppControllerRoleGranted, error) {
	event := new(AppControllerRoleGranted)
	if err := _AppController.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppControllerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AppController contract.
type AppControllerRoleRevokedIterator struct {
	Event *AppControllerRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AppControllerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppControllerRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AppControllerRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AppControllerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppControllerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppControllerRoleRevoked represents a RoleRevoked event raised by the AppController contract.
type AppControllerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppController *AppControllerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AppControllerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AppController.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AppControllerRoleRevokedIterator{contract: _AppController.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppController *AppControllerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AppControllerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AppController.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppControllerRoleRevoked)
				if err := _AppController.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppController *AppControllerFilterer) ParseRoleRevoked(log types.Log) (*AppControllerRoleRevoked, error) {
	event := new(AppControllerRoleRevoked)
	if err := _AppController.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
