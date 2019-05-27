package business

type VMHosts struct {
	Message    string      `json:"message"`
	Code       int         `json:"code"`
	Data       VMHostsData `json:"data"`
	Request_id string      `json:"request_id"`
	Result     bool        `json:"result"`
}

type VMHostsData struct {
	Count int      `json:"count"`
	Info  []VMinfo `json:"info"`
}

type VMinfo struct {
	Biz    []string   `json:"biz"`
	Host   HostDetail `json:"host"`
	Module []string   `json:"module"`
	Set    []string   `json:"set"`
}

type HostDetail struct {
	//bk_property_id  string
	bk_host_innerip *string
	bk_host_outerip *string
	//bk_bak_operator string
	bk_asset_id *string
	//operator        string
	bk_sn      *string
	bk_comment *string
	//bk_service_term string
	//bk_sla          string
	//////////////bk_cloud_id *[]Cloud
	//bk_state_name string
	//bk_province_name string
	//bk_isp_name      string
	bk_host_name  *string
	bk_os_type    *string
	bk_os_name    *string
	bk_os_version *string
	bk_cpu        *int
	bk_os_bit     *string
	bk_cpu_mhz    *int
	bk_cpu_module *string
	bk_mem        *int
	bk_disk       *int
	bk_mac        *string
	bk_outer_mac  *string
	create_time   *string
	//import_from      string
	lease          *string
	nas            *string
	host_type      *string
	logic_disk_num *string
	v_mem          *string
	subnet_mask    *string
	dns_serv       *string
	gateway        *string
	//sys_install_time string
	time_zone_set    *string
	ntp_srv          *string
	last_reboot_time *string
	manufacturer     *string
	host_model       *string
	net_card_rate    *string
	////////////running_on_vp    *VI
	////////////running_on_srv   *[]PHost
	bk_host_id *int
}
