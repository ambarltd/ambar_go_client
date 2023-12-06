# MysqlDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | A case sensitive string for the user Ambar should use to connect to your MySQL database. | 
**Password** | **string** | A case sensitive string for the user Ambar should use to connect to your MySQL database. | 
**Hostname** | **string** | A case insensitive string for the host on which your MySQL database is running and which Ambar can use to reach the database. | 
**HostPort** | **string** | The port number passed as a string which Ambar can use to connect to your MySQL database instance. | 
**DatabaseName** | **string** | The case sensitive string name of the database on your database host. | 
**TableName** | **string** | The case sensitive string name of the table the DataSource should read. | 
**GloballyUniqueColumnName** | **string** | The name of a column where the value for any record is globally unique. | 
**IncrementingColumnName** | **string** | The name of a column which monotonically increases on database writes. | 
**PartitioningColumnName** | **string** | A case sensitive string for the name of the column in your table Ambar can partition on.  Note that partition keys must be unique to a given sequence of records. | 
**AdditionalColumns** | **string** | A case sensitive, comma separated list string of additional columns to be read from  the table into Ambar. | 
**BinLogReplicationServerId** | **string** | The server_id value used when starting the MySQL server. See MySQL docs for more information about this value and its defaults when not configured. | 
**TlsTerminationOverrideHost** | **string** | The hostname of the server responsible for terminating TLS connections for your server,  for example if your server is behind a load balancer. | 

## Methods

### NewMysqlDataSource

`func NewMysqlDataSource(username string, password string, hostname string, hostPort string, databaseName string, tableName string, globallyUniqueColumnName string, incrementingColumnName string, partitioningColumnName string, additionalColumns string, binLogReplicationServerId string, tlsTerminationOverrideHost string, ) *MysqlDataSource`

NewMysqlDataSource instantiates a new MysqlDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlDataSourceWithDefaults

`func NewMysqlDataSourceWithDefaults() *MysqlDataSource`

NewMysqlDataSourceWithDefaults instantiates a new MysqlDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *MysqlDataSource) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MysqlDataSource) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MysqlDataSource) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *MysqlDataSource) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MysqlDataSource) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MysqlDataSource) SetPassword(v string)`

SetPassword sets Password field to given value.


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


### GetGloballyUniqueColumnName

`func (o *MysqlDataSource) GetGloballyUniqueColumnName() string`

GetGloballyUniqueColumnName returns the GloballyUniqueColumnName field if non-nil, zero value otherwise.

### GetGloballyUniqueColumnNameOk

`func (o *MysqlDataSource) GetGloballyUniqueColumnNameOk() (*string, bool)`

GetGloballyUniqueColumnNameOk returns a tuple with the GloballyUniqueColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGloballyUniqueColumnName

`func (o *MysqlDataSource) SetGloballyUniqueColumnName(v string)`

SetGloballyUniqueColumnName sets GloballyUniqueColumnName field to given value.


### GetIncrementingColumnName

`func (o *MysqlDataSource) GetIncrementingColumnName() string`

GetIncrementingColumnName returns the IncrementingColumnName field if non-nil, zero value otherwise.

### GetIncrementingColumnNameOk

`func (o *MysqlDataSource) GetIncrementingColumnNameOk() (*string, bool)`

GetIncrementingColumnNameOk returns a tuple with the IncrementingColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementingColumnName

`func (o *MysqlDataSource) SetIncrementingColumnName(v string)`

SetIncrementingColumnName sets IncrementingColumnName field to given value.


### GetPartitioningColumnName

`func (o *MysqlDataSource) GetPartitioningColumnName() string`

GetPartitioningColumnName returns the PartitioningColumnName field if non-nil, zero value otherwise.

### GetPartitioningColumnNameOk

`func (o *MysqlDataSource) GetPartitioningColumnNameOk() (*string, bool)`

GetPartitioningColumnNameOk returns a tuple with the PartitioningColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioningColumnName

`func (o *MysqlDataSource) SetPartitioningColumnName(v string)`

SetPartitioningColumnName sets PartitioningColumnName field to given value.


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


