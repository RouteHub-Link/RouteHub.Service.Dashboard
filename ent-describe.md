Domain:
	+--------+--------------------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	| Field  |        Type        | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |        StructTag        | Validators | Comment |
	+--------+--------------------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	| id     | schema.DomainID    | false  | false    | false    | true    | false         | false     | json:"id,omitempty"     |          0 |         |
	| name   | string             | false  | false    | false    | false   | false         | false     | json:"name,omitempty"   |          1 |         |
	| url    | string             | true   | false    | false    | false   | false         | false     | json:"url,omitempty"    |          1 |         |
	| status | domain.DomainState | false  | false    | false    | false   | false         | false     | json:"status,omitempty" |          0 |         |
	+--------+--------------------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	|     Edge     |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	| hub          | Hub          | true    | domain  | O2O      | true   | true     |         |
	| organization | Organization | true    | domains | M2O      | true   | false    |         |
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	
Hub:
	+---------------------+-----------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------------+------------+---------+
	|        Field        |         Type          | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |              StructTag               | Validators | Comment |
	+---------------------+-----------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------------+------------+---------+
	| id                  | schema.HubID          | false  | false    | false    | true    | false         | false     | json:"id,omitempty"                  |          0 |         |
	| name                | string                | false  | false    | false    | false   | false         | false     | json:"name,omitempty"                |          1 |         |
	| slug                | string                | false  | false    | false    | false   | false         | false     | json:"slug,omitempty"                |          1 |         |
	| Hub_details         | types.HubDetails      | false  | true     | false    | false   | false         | false     | json:"Hub_details,omitempty"         |          0 |         |
	| TCPAddress          | string                | false  | false    | false    | false   | false         | false     | json:"TCPAddress,omitempty"          |          1 |         |
	| status              | enums.StatusState     | false  | false    | false    | false   | false         | false     | json:"status,omitempty"              |          0 |         |
	| default_redirection | hub.RedirectionOption | false  | false    | false    | false   | false         | false     | json:"default_redirection,omitempty" |          0 |         |
	+---------------------+-----------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------------+------------+---------+
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	|     Edge     |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	| domain       | Domain       | false   |         | O2O      | true   | false    |         |
	| organization | Organization | true    | hubs    | M2O      | true   | false    |         |
	| links        | Link         | false   |         | O2M      | false  | true     |         |
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	
Link:
	+--------------+-------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------+------------+---------+
	|    Field     |       Type        | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag           | Validators | Comment |
	+--------------+-------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------+------------+---------+
	| id           | schema.LinkID     | false  | false    | false    | true    | false         | false     | json:"id,omitempty"           |          0 |         |
	| target       | string            | false  | false    | false    | false   | false         | false     | json:"target,omitempty"       |          1 |         |
	| path         | string            | true   | false    | false    | false   | false         | false     | json:"path,omitempty"         |          1 |         |
	| link_content | types.LinkContent | false  | true     | false    | false   | false         | false     | json:"link_content,omitempty" |          0 |         |
	| status       | enums.StatusState | false  | false    | false    | false   | false         | false     | json:"status,omitempty"       |          0 |         |
	+--------------+-------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------+------------+---------+
	+------+------+---------+---------+----------+--------+----------+---------+
	| Edge | Type | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+------+------+---------+---------+----------+--------+----------+---------+
	| hub  | Hub  | true    | links   | M2O      | true   | false    |         |
	+------+------+---------+---------+----------+--------+----------+---------+
	
Organization:
	+---------------+-----------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	|     Field     |         Type          | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag            | Validators | Comment |
	+---------------+-----------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	| id            | schema.OrganizationID | false  | false    | false    | true    | false         | false     | json:"id,omitempty"            |          0 |         |
	| name          | string                | false  | false    | false    | false   | false         | false     | json:"name,omitempty"          |          1 |         |
	| website       | string                | false  | true     | false    | false   | false         | false     | json:"website,omitempty"       |          0 |         |
	| description   | string                | false  | true     | false    | false   | false         | false     | json:"description,omitempty"   |          0 |         |
	| locagtion     | string                | false  | true     | false    | false   | false         | false     | json:"locagtion,omitempty"     |          0 |         |
	| social_medias | types.SocialMedias    | false  | true     | false    | false   | false         | false     | json:"social_medias,omitempty" |          0 |         |
	+---------------+-----------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	+---------+--------+---------+---------+----------+--------+----------+---------+
	|  Edge   |  Type  | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------+--------+---------+---------+----------+--------+----------+---------+
	| domains | Domain | false   |         | O2M      | false  | true     |         |
	| hubs    | Hub    | false   |         | O2M      | false  | true     |         |
	+---------+--------+---------+---------+----------+--------+----------+---------+
	
