# MysqlDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | A case insensitive string for the host on which your MySQL database is running and which Ambar can use to reach the database. | 
**HostPort** | **string** | The port number passed as a string which Ambar can use to connect to your MySQL database instance. | 
**DatabaseName** | **string** | The case sensitive string name of the database on your database host. | 
**TableName** | **string** | The case sensitive string name of the table the DataSource should read. | 
**GloballyUniqueColumn** | **string** | The name of a column where the value for any record is globally unique. | 
**SerialColumn** | **string** | The name of a column which monotonically increases on database writes. | 
**PartitioningColumn** | **string** | A case sensitive string for the name of the column in your table Ambar can partition on.  Note that partition keys must be unique to a given sequence of records. | 
**AdditionalColumns** | **string** | A case sensitive, comma separated list string of additional columns to be read from  the table into Ambar. | 
**BinLogReplicationServerId** | **string** | The server_id value used when starting the MySQL server. See MySQL docs for more information about this value and its defaults when not configured. | 
**TlsTerminationOverrideHost** | **string** | The hostname of the server responsible for terminating TLS connections for your server,  for example if your server is behind a load balancer. | 

## Methods

### NewMysqlDataSource

`func NewMysqlDataSource(hostname string, hostPort string, databaseName string, tableName string, globallyUniqueColumn string, serialColumn string, partitioningColumn string, additionalColumns string, binLogReplicationServerId string, tlsTerminationOverrideHost string, ) *MysqlDataSource`

NewMysqlDataSource instantiates a new MysqlDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlDataSourceWithDefaults

`func NewMysqlDataSourceWithDefaults() *MysqlDataSource`

NewMysqlDataSourceWithDefaults instantiates a new MysqlDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *MysqlDataSource) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *MysqlDataSource) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *MysqlDataSource) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetHostPort

`func (o *MysqlDataSource) GetHostPort() string`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *MysqlDataSource) GetHostPortOk() (*string, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *MysqlDataSource) SetHostPort(v string)`

SetHostPort sets HostPort field to given value.


### GetDatabaseName

`func (o *MysqlDataSource) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *MysqlDataSource) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *MysqlDataSource) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetTableName

`func (o *MysqlDataSource) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *MysqlDataSource) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *MysqlDataSource) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetGloballyUniqueColumn

`func (o *MysqlDataSource) GetGloballyUniqueColumn() string`

GetGloballyUniqueColumn returns the GloballyUniqueColumn field if non-nil, zero value otherwise.

### GetGloballyUniqueColumnOk

`func (o *MysqlDataSource) GetGloballyUniqueColumnOk() (*string, bool)`

GetGloballyUniqueColumnOk returns a tuple with the GloballyUniqueColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGloballyUniqueColumn

`func (o *MysqlDataSource) SetGloballyUniqueColumn(v string)`

SetGloballyUniqueColumn sets GloballyUniqueColumn field to given value.


### GetSerialColumn

`func (o *MysqlDataSource) GetSerialColumn() string`

GetSerialColumn returns the SerialColumn field if non-nil, zero value otherwise.

### GetSerialColumnOk

`func (o *MysqlDataSource) GetSerialColumnOk() (*string, bool)`

GetSerialColumnOk returns a tuple with the SerialColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialColumn

`func (o *MysqlDataSource) SetSerialColumn(v string)`

SetSerialColumn sets SerialColumn field to given value.


### GetPartitioningColumn

`func (o *MysqlDataSource) GetPartitioningColumn() string`

GetPartitioningColumn returns the PartitioningColumn field if non-nil, zero value otherwise.

### GetPartitioningColumnOk

`func (o *MysqlDataSource) GetPartitioningColumnOk() (*string, bool)`

GetPartitioningColumnOk returns a tuple with the PartitioningColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioningColumn

`func (o *MysqlDataSource) SetPartitioningColumn(v string)`

SetPartitioningColumn sets PartitioningColumn field to given value.


### GetAdditionalColumns

`func (o *MysqlDataSource) GetAdditionalColumns() string`

GetAdditionalColumns returns the AdditionalColumns field if non-nil, zero value otherwise.

### GetAdditionalColumnsOk

`func (o *MysqlDataSource) GetAdditionalColumnsOk() (*string, bool)`

GetAdditionalColumnsOk returns a tuple with the AdditionalColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalColumns

`func (o *MysqlDataSource) SetAdditionalColumns(v string)`

SetAdditionalColumns sets AdditionalColumns field to given value.


### GetBinLogReplicationServerId

`func (o *MysqlDataSource) GetBinLogReplicationServerId() string`

GetBinLogReplicationServerId returns the BinLogReplicationServerId field if non-nil, zero value otherwise.

### GetBinLogReplicationServerIdOk

`func (o *MysqlDataSource) GetBinLogReplicationServerIdOk() (*string, bool)`

GetBinLogReplicationServerIdOk returns a tuple with the BinLogReplicationServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinLogReplicationServerId

`func (o *MysqlDataSource) SetBinLogReplicationServerId(v string)`

SetBinLogReplicationServerId sets BinLogReplicationServerId field to given value.


### GetTlsTerminationOverrideHost

`func (o *MysqlDataSource) GetTlsTerminationOverrideHost() string`

GetTlsTerminationOverrideHost returns the TlsTerminationOverrideHost field if non-nil, zero value otherwise.

### GetTlsTerminationOverrideHostOk

`func (o *MysqlDataSource) GetTlsTerminationOverrideHostOk() (*string, bool)`

GetTlsTerminationOverrideHostOk returns a tuple with the TlsTerminationOverrideHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTerminationOverrideHost

`func (o *MysqlDataSource) SetTlsTerminationOverrideHost(v string)`

SetTlsTerminationOverrideHost sets TlsTerminationOverrideHost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


