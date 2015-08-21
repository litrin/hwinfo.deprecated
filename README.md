
# common
    import "github.com/mickep76/hwinfo/common"






## func ExecCmdFields
``` go
func ExecCmdFields(cmd string, args []string, del string, fields []string) (map[string]string, error)
```
ExecCmdFields returns fields from command output.


## func LoadFileFields
``` go
func LoadFileFields(fn string, del string, fields []string) (map[string]string, error)
```
LoadFileFields returns fields from file.









- - -

# cpu
    import "github.com/mickep76/hwinfo/cpu"







## type Info
``` go
type Info struct {
    Model          string `json:"model"`
    Flags          string `json:"flags"`
    Logical        int    `json:"logical"`
    Physical       int    `json:"physical"`
    Sockets        int    `json:"sockets"`
    CoresPerSocket int    `json:"cores_per_socket"`
    ThreadsPerCore int    `json:"threads_per_core"`
}
```
Info structure for information about a systems CPU(s).









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a systems CPU(s).










- - -
