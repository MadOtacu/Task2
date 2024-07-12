module example.com

go 1.22.4

replace example.com/server => ../server

replace example.com/fileSys => ../fileSys

replace example.com/serverOutput => ../serverOutput

replace example.com/fileSys => ../fileSys

require gopkg.in/ini.v1 v1.67.0

require github.com/stretchr/testify v1.9.0 // indirect

replace example.com/sorting => ../sorting
