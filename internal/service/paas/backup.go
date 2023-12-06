package paas

import (
	"context"
	"github.com/aws/aws-sdk-go/service/paas"
	"github.com/hashicorp/terraform-provider-aws/internal/experimental/nullable"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func ResourceBackup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBackupCreate,
		ReadContext:   resourceBackupRead,
		UpdateContext: resourceBackupUpdate,
		DeleteContext: resourceBackupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"backup_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logfile": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"enable_deletion_protection": {
				Type:     nullable.TypeNullableBool,
				Optional: true,
			},
			"force_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"protected": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"service_class": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_deleted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceBackupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("backup_id").(string)
	d.SetId(id)

	return resourceBackupUpdate(ctx, d, meta)
}

func resourceBackupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).PaaSConn
	id := d.Id()

	backup, err := FindBackupById(conn, id)

	if err != nil {
		return diag.Errorf("error reading PaaS Backup (%s): %s", id, err)
	}

	d.Set("enable_deletion_protection", backup.Protected)
	d.Set("protected", backup.Protected)
	d.Set("service_class", backup.ServiceClass)
	d.Set("service_deleted", backup.ServiceDeleted)
	d.Set("service_id", backup.ServiceId)
	d.Set("service_name", backup.ServiceName)
	d.Set("service_type", backup.ServiceType)
	d.Set("status", backup.Status)

	if backup.Time != nil {
		d.Set("time", time.Unix(aws.Int64Value(backup.Time), 0).Format(time.RFC3339))
	}

	d.Set("databases", flattenDatabaseBackups(backup.Databases))

	return nil
}

func resourceBackupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).PaaSConn
	id := d.Id()

	if v, null, err := nullable.Bool(d.Get("enable_deletion_protection").(string)).Value(); !null {
		if err != nil {
			return diag.Errorf("error reading `enable_deletion_protection` from configuration: %s", err)
		}

		_, err := updateBackupProtection(conn, id, v)

		if err != nil {
			return diag.Errorf("error modifying PaaS Service Backup (%s): %s", id, err)
		}
	}

	return resourceBackupRead(ctx, d, meta)
}

func resourceBackupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).PaaSConn
	id := d.Id()

	if d.Get("force_delete").(bool) {
		_, err := updateBackupProtection(conn, id, false)

		if err != nil {
			return diag.Errorf("error modifying PaaS Service Backup (%s): %s", id, err)
		}

		input := &paas.DeleteBackupsInput{
			ServiceId: aws.String(d.Get("service_id").(string)),
			BackupIds: []*string{aws.String(id)},
		}

		log.Printf("[DEBUG] Deleting PaaS Service Backup: %s", input)
		_, err = conn.DeleteBackups(input)

		if err != nil {
			return diag.Errorf("error deleting PaaS Service Backup (%s): %s", id, err)
		}
	}

	log.Printf("[WARN] PaaS Service Backup (%s) was not deleted, removing from state", id)

	return nil
}

func updateBackupProtection(conn *paas.PaaS, id string, protected bool) (*paas.Backup, error) { //nolint:unparam
	input := &paas.ModifyBackupInput{
		BackupId:  aws.String(id),
		Protected: aws.Bool(protected),
	}
	log.Printf("[DEBUG] Modifying PaaS Service Backup: %s", input)

	res, err := conn.ModifyBackup(input)

	if err != nil {
		return nil, err
	}

	return res.Backup, err
}
