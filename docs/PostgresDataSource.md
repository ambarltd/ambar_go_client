# PostgresDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | A case insensitive string for the host on which your Postgres database is running and which Ambar can use to reach the database. | 
**HostPort** | **string** | The port number passed as a string which Ambar can use to connect to your Postgres database instance. | 
**DatabaseName** | **string** | The case sensitive string name of the database on your database host. | 
**TableName** | **string** | The case sensitive string name of the table the DataSource should read. | 
**PublicationName** | **string** | The named publication Ambar should use to indicate to Postgres what tables Ambar will be replicating from.  The publication name must be distinct from the name of any existing publication in the current database. See Postgres documentation for your specific version of Postgres for more information on Postgres publications. | 
**AdditionalColumns** | **string** | A case sensitive, comma separated list string of columns which Ambar should read from the database. The ordering should be consistent with the table from which Ambar will read. | 

## Methods

### NewPostgresDataSource

`func NewPostgresDataSource(hostname string, hostPort string, databaseName string, tableName string, publicationName string, additionalColumns string, ) *PostgresDataSource`

NewPostgresDataSource instantiates a new PostgresDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresDataSourceWithDefaults

`func NewPostgresDataSourceWithDefaults() *PostgresDataSource`

NewPostgresDataSourceWithDefaults instantiates a new PostgresDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *PostgresDataSource) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PostgresDataSource) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PostgresDataSource) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetHostPort

`func (o *PostgresDataSource) GetHostPort() string`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *PostgresDataSource) GetHostPortOk() (*string, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *PostgresDataSource) SetHostPort(v string)`

SetHostPort sets HostPort field to given value.


### GetDatabaseName

`func (o *PostgresDataSource) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *PostgresDataSource) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *PostgresDataSource) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetTableName

`func (o *PostgresDataSource) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *PostgresDataSource) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *PostgresDataSource) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetPublicationName

`func (o *PostgresDataSource) GetPublicationName() string`

GetPublicationName returns the PublicationName field if non-nil, zero value otherwise.

### GetPublicationNameOk

`func (o *PostgresDataSource) GetPublicationNameOk() (*string, bool)`

GetPublicationNameOk returns a tuple with the PublicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationName

`func (o *PostgresDataSource) SetPublicationName(v string)`

SetPublicationName sets PublicationName field to given value.


### GetAdditionalColumns

`func (o *PostgresDataSource) GetAdditionalColumns() string`

GetAdditionalColumns returns the AdditionalColumns field if non-nil, zero value otherwise.

### GetAdditionalColumnsOk

`func (o *PostgresDataSource) GetAdditionalColumnsOk() (*string, bool)`

GetAdditionalColumnsOk returns a tuple with the AdditionalColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalColumns

`func (o *PostgresDataSource) SetAdditionalColumns(v string)`

SetAdditionalColumns sets AdditionalColumns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


