package yandex

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1"
	"github.com/yandex-cloud/go-sdk/sdkresolvers"
)

func dataSourceYandexMDBPostgreSQLCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceYandexMDBPostgreSQLClusterRead,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     dataSourceYandexMDBPostgreSQLClusterConfigBlock(),
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"folder_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"health": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     dataSourceYandexMDBPostgreSQLClusterHostBlock(),
			},
			"labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_group_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Computed: true,
			},
			"maintenance_window": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     dataSourceYandexMDBPostgreSQLClusterMaintenanceWindowBlock(),
			},
			"deletion_protection": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"host_group_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Computed: true,
			},
			"database": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      mysqlDatabaseHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"owner": {
							Type:     schema.TypeString,
							ForceNew: true,
							Required: true,
						},
						"lc_collate": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Default:  "C",
						},
						"lc_type": {
							Type:     schema.TypeString,
							ForceNew: true,
							Optional: true,
							Default:  "C",
						},
						"template_db": {
							Type:     schema.TypeString,
							ForceNew: true,
							Optional: true,
						},
						"extension": {
							Type:     schema.TypeSet,
							Set:      pgExtensionHash,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"version": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"user": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"login": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"grants": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},
						// TODO change to permissions
						"permission": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      pgUserPermissionHash,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"database_name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"conn_limit": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"settings": {
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: generateMapSchemaDiffSuppressFunc(mdbPGUserSettingsFieldsInfo),
							ValidateFunc:     generateMapSchemaValidateFunc(mdbPGUserSettingsFieldsInfo),
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceYandexMDBPostgreSQLClusterConfigBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_lens": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"web_sql": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"serverless": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"data_transfer": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"autofailover": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"backup_window_start": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hours": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"minutes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"backup_retain_period_days": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"performance_diagnostics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sessions_sampling_interval": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"statements_sampling_interval": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"disk_size_autoscaling": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_size_limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"planned_usage_threshold": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"emergency_usage_threshold": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"pooler_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pool_discard": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"pooling_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"disk_type_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_preset_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"postgresql_config": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: generateMapSchemaDiffSuppressFunc(mdbPGSettingsFieldsInfo),
				ValidateFunc:     generateMapSchemaValidateFunc(mdbPGSettingsFieldsInfo),
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceYandexMDBPostgreSQLClusterHostBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"assign_public_ip": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:       schema.TypeInt,
				Computed:   true,
				Deprecated: "The field has not affected anything. You can safely delete it.",
			},
		},
	}
}

func dataSourceYandexMDBPostgreSQLClusterMaintenanceWindowBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"day": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hour": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceYandexMDBPostgreSQLClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	ctx := config.Context()

	err := checkOneOf(d, "cluster_id", "name")
	if err != nil {
		return err
	}

	clusterID := d.Get("cluster_id").(string)
	_, clusterNameOk := d.GetOk("name")

	if clusterNameOk {
		clusterID, err = resolveObjectID(ctx, config, d, sdkresolvers.PostgreSQLClusterResolver)
		if err != nil {
			return fmt.Errorf("failed to resolve data source PostgreSQL Cluster by name: %v", err)
		}
	}

	cluster, err := config.sdk.MDB().PostgreSQL().Cluster().Get(ctx, &postgresql.GetClusterRequest{
		ClusterId: clusterID,
	})
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Cluster %q", clusterID))
	}

	databases, err := listPGDatabases(ctx, config, clusterID)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Cluster %q", clusterID))
	}
	if err := d.Set("database", flattenPGDatabases(databases)); err != nil {
		return err
	}

	passwords := pgUsersPasswords(make([]*postgresql.UserSpec, 0))
	users, err := listPGUsers(ctx, config, clusterID)
	if err != nil {
		return err
	}
	fUsers, err := flattenPGUsers(users, passwords, mdbPGUserSettingsFieldsInfo)
	if err != nil {
		return err
	}
	if err := d.Set("user", fUsers); err != nil {
		return err
	}

	pgClusterConfig, err := flattenPGClusterConfig(cluster.Config)
	if err != nil {
		return err
	}
	if err := d.Set("config", pgClusterConfig); err != nil {
		return err
	}

	hosts, err := retryListPGHostsWrapper(ctx, config, clusterID)
	if err != nil {
		return err
	}

	orderedHostInfos, err := flattenPGHostsInfo(d, hosts)
	if err != nil {
		return err
	}

	hs := flattenPGHostsFromHostInfos(d, orderedHostInfos, true)
	if err := d.Set("host", hs); err != nil {
		return err
	}

	if err := d.Set("labels", cluster.Labels); err != nil {
		return err
	}

	if err := d.Set("security_group_ids", cluster.SecurityGroupIds); err != nil {
		return err
	}

	maintenanceWindow, err := flattenPGMaintenanceWindow(cluster.MaintenanceWindow)
	if err != nil {
		return err
	}

	if err := d.Set("maintenance_window", maintenanceWindow); err != nil {
		return err
	}

	if err = d.Set("host_group_ids", cluster.HostGroupIds); err != nil {
		return err
	}

	d.Set("created_at", getTimestamp(cluster.CreatedAt))
	d.Set("cluster_id", cluster.Id)
	d.Set("name", cluster.Name)
	d.Set("folder_id", cluster.FolderId)
	d.Set("network_id", cluster.NetworkId)
	d.Set("environment", cluster.GetEnvironment().String())
	d.Set("health", cluster.GetHealth().String())
	d.Set("status", cluster.GetStatus().String())
	d.Set("description", cluster.Description)
	d.Set("deletion_protection", cluster.DeletionProtection)

	d.SetId(cluster.Id)
	return nil
}
