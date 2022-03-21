# easy-sync-cli
this cli tool allows you to synch users (members) between various online protal

# WIP: installation
for now go and GNU make need to be installed
- clone the repo `https://github.com/funkymcb/easy-sync.git`
- run `make build`
- add ./out/easy-sync binary to your PATH

a simpler installation will be added in the future

# offered features
- read users from csv file
- create valid json from csv
- synch from csv to easyverein.com
- synch between easyverein and wordpress

# roadmap
- [x] create basic cli handling
- [x] implement csv file reader
- [x] implement csv to json decoding
    - TODO fix umlauts
- [ ] synch json with easyverein
- [ ] synch json with wordpress
- [ ] synch wordpress with easyverein
