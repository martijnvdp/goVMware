package functions

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type ByName []mo.VirtualMachine

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }

//ListVMS returns a list of vms
func ListVMS() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	// Parse URL from string
	u, err := url.Parse(*urlFlag)
	if err != nil {
		exit(err)
	}

	// Override username and/or password as required
	processOverride(u)

	// Connect and log in to ESX or vCenter
	c, err := govmomi.NewClient(ctx, u, *insecureFlag)
	if err != nil {
		exit(err)
	}

	f := find.NewFinder(c.Client, true)

	// Find one and only datacenter
	dc, err := f.DefaultDatacenter(ctx)
	if err != nil {
		exit(err)
	}

	// Make future calls local to this datacenter
	f.SetDatacenter(dc)

	// Find virtual machines in datacenter
	vms, err := f.VirtualMachineList(ctx, "*")
	if err != nil {
		exit(err)
	}

	pc := property.DefaultCollector(c.Client)

	// Convert datastores into list of references
	var refs []types.ManagedObjectReference
	for _, vm := range vms {
		refs = append(refs, vm.Reference())
	}

	// Retrieve name property for all vms
	var vmt []mo.VirtualMachine
	err = pc.Retrieve(ctx, refs, []string{"name"}, &vmt)
	if err != nil {
		exit(err)
	}

	// Print name per virtual machine
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	fmt.Println("Virtual machines found:", len(vmt))
	sort.Sort(ByName(vmt))
	for _, vm := range vmt {
		fmt.Fprintf(tw, "%s\n", vm.Name)
	}
	tw.Flush()
}
