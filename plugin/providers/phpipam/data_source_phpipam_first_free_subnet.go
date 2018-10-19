package phpipam

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourcePHPIPAMFirstFreeSubnet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePHPIPAMFirstFreeSubnetRead,
		Schema: map[string]*schema.Schema{
			"subnet_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"mask": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"subnet_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePHPIPAMFirstFreeSubnetRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*ProviderPHPIPAMClient).subnetsController
	out, err := c.GetFirstFreeSubnet(d.Get("subnet_id").(int), d.Get("mask").(int))
	if err != nil {
		return err
	}
	if out == "" {
		return fmt.Errorf("Subnet has no free subnet addresses with mask %d", d.Get("mask").(int))
	}

	d.SetId(out)
	d.Set("subnet_cidr", out)

	return nil
}
