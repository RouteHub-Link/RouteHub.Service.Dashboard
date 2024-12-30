Domain:
	+------------+--------------------+--------+----------+----------+---------+---------------+-----------+-----------------------------+------------+---------+
	|   Field    |        Type        | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |          StructTag          | Validators | Comment |
	+------------+--------------------+--------+----------+----------+---------+---------------+-----------+-----------------------------+------------+---------+
	| id         | mixin.ID           | false  | false    | false    | true    | false         | false     | json:"id,omitempty"         |          0 |         |
	| name       | string             | false  | false    | false    | false   | false         | false     | json:"name,omitempty"       |          1 |         |
	| url        | string             | true   | false    | false    | false   | false         | false     | json:"url,omitempty"        |          1 |         |
	| status     | domain.DomainState | false  | false    | false    | false   | false         | false     | json:"status,omitempty"     |          0 |         |
	| created_at | time.Time          | false  | false    | false    | true    | false         | false     | json:"created_at,omitempty" |          0 |         |
	+------------+--------------------+--------+----------+----------+---------+---------------+-----------+-----------------------------+------------+---------+
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	|     Edge     |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	| organization | Organization | true    | domains | M2O      | true   | false    |         |
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	
Hub:
	+---------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------------+------------+---------+
	|        Field        |          Type           | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |              StructTag               | Validators | Comment |
	+---------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------------+------------+---------+
	| id                  | mixin.ID                | false  | false    | false    | true    | false         | false     | json:"id,omitempty"                  |          0 |         |
	| name                | string                  | false  | false    | false    | false   | false         | false     | json:"name,omitempty"                |          1 |         |
	| slug                | string                  | false  | false    | false    | false   | false         | false     | json:"slug,omitempty"                |          1 |         |
	| hub_details         | types.HubDetails        | false  | true     | false    | false   | false         | false     | json:"hub_details,omitempty"         |          0 |         |
	| tcp_address         | string                  | false  | false    | false    | false   | false         | false     | json:"tcp_address,omitempty"         |          1 |         |
	| status              | enums.StatusState       | false  | false    | false    | false   | false         | false     | json:"status,omitempty"              |          0 |         |
	| default_redirection | enums.RedirectionChoice | false  | false    | false    | false   | false         | false     | json:"default_redirection,omitempty" |          0 |         |
	| created_at          | time.Time               | false  | false    | false    | true    | false         | false     | json:"created_at,omitempty"          |          0 |         |
	+---------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------------+------------+---------+
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	|     Edge     |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	| domain       | Domain       | false   |         | M2O      | true   | false    |         |
	| organization | Organization | true    | hubs    | M2O      | true   | false    |         |
	| links        | Link         | false   |         | O2M      | false  | true     |         |
	| pages        | Page         | false   |         | O2M      | false  | true     |         |
	+--------------+--------------+---------+---------+----------+--------+----------+---------+
	
Link:
	+--------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	|       Field        |          Type           | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |              StructTag              | Validators | Comment |
	+--------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	| id                 | mixin.ID                | false  | false    | false    | true    | false         | false     | json:"id,omitempty"                 |          0 |         |
	| target             | string                  | false  | false    | false    | false   | false         | false     | json:"target,omitempty"             |          1 |         |
	| path               | string                  | true   | false    | false    | false   | false         | false     | json:"path,omitempty"               |          1 |         |
	| link_content       | *types.LinkContent      | false  | true     | false    | true    | false         | false     | json:"link_content,omitempty"       |          0 |         |
	| status             | enums.StatusState       | false  | false    | false    | false   | false         | false     | json:"status,omitempty"             |          0 |         |
	| redirection_choice | enums.RedirectionChoice | false  | false    | false    | false   | false         | false     | json:"redirection_choice,omitempty" |          0 |         |
	| created_at         | time.Time               | false  | false    | false    | true    | false         | false     | json:"created_at,omitempty"         |          0 |         |
	+--------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	+------+------+---------+---------+----------+--------+----------+---------+
	| Edge | Type | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+------+------+---------+---------+----------+--------+----------+---------+
	| hub  | Hub  | true    | links   | M2O      | true   | false    |         |
	+------+------+---------+---------+----------+--------+----------+---------+
	
Organization:
	+---------------+--------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	|     Field     |        Type        | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag            | Validators | Comment |
	+---------------+--------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	| id            | mixin.ID           | false  | false    | false    | true    | false         | false     | json:"id,omitempty"            |          0 |         |
	| name          | string             | false  | false    | false    | false   | false         | false     | json:"name,omitempty"          |          1 |         |
	| website       | string             | false  | true     | false    | false   | false         | false     | json:"website,omitempty"       |          0 |         |
	| description   | string             | false  | true     | false    | false   | false         | false     | json:"description,omitempty"   |          0 |         |
	| location      | string             | false  | true     | false    | false   | false         | false     | json:"location,omitempty"      |          0 |         |
	| social_medias | types.SocialMedias | false  | true     | false    | false   | false         | false     | json:"social_medias,omitempty" |          0 |         |
	| created_at    | time.Time          | false  | false    | false    | true    | false         | false     | json:"created_at,omitempty"    |          0 |         |
	+---------------+--------------------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	+---------+--------+---------+---------+----------+--------+----------+---------+
	|  Edge   |  Type  | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------+--------+---------+---------+----------+--------+----------+---------+
	| domains | Domain | false   |         | O2M      | false  | true     |         |
	| hubs    | Hub    | false   |         | O2M      | false  | true     |         |
	| persons | Person | false   |         | M2M      | false  | true     |         |
	+---------+--------+---------+---------+----------+--------+----------+---------+
	
Page:
	+-------------------+-----------------------+--------+----------+----------+---------+---------------+-----------+------------------------------------+------------+---------+
	|       Field       |         Type          | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |             StructTag              | Validators | Comment |
	+-------------------+-----------------------+--------+----------+----------+---------+---------------+-----------+------------------------------------+------------+---------+
	| id                | mixin.ID              | false  | false    | false    | true    | false         | false     | json:"id,omitempty"                |          0 |         |
	| name              | string                | false  | false    | false    | false   | false         | false     | json:"name,omitempty"              |          1 |         |
	| slug              | string                | false  | false    | false    | false   | false         | false     | json:"slug,omitempty"              |          1 |         |
	| page_description  | string                | false  | true     | false    | false   | false         | false     | json:"page_description,omitempty"  |          0 |         |
	| page_content_json | string                | false  | true     | false    | false   | false         | false     | json:"page_content_json,omitempty" |          0 |         |
	| page_content_html | string                | false  | true     | false    | false   | false         | false     | json:"page_content_html,omitempty" |          0 |         |
	| status            | enums.StatusState     | false  | false    | false    | false   | false         | false     | json:"status,omitempty"            |          0 |         |
	| meta_description  | types.MetaDescription | false  | true     | false    | false   | false         | false     | json:"meta_description,omitempty"  |          0 |         |
	| created_at        | time.Time             | false  | false    | false    | true    | false         | false     | json:"created_at,omitempty"        |          0 |         |
	+-------------------+-----------------------+--------+----------+----------+---------+---------------+-----------+------------------------------------+------------+---------+
	+------+------+---------+---------+----------+--------+----------+---------+
	| Edge | Type | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+------+------+---------+---------+----------+--------+----------+---------+
	| hub  | Hub  | true    | pages   | M2O      | true   | false    |         |
	+------+------+---------+---------+----------+--------+----------+---------+
	
Person:
	+------------+---------------+--------+----------+----------+---------+---------------+-----------+-----------------------------+------------+---------+
	|   Field    |     Type      | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |          StructTag          | Validators | Comment |
	+------------+---------------+--------+----------+----------+---------+---------------+-----------+-----------------------------+------------+---------+
	| id         | mixin.ID      | false  | false    | false    | true    | false         | false     | json:"id,omitempty"         |          0 |         |
	| subject_id | string        | true   | false    | false    | false   | false         | false     | json:"subject_id,omitempty" |          1 |         |
	| user_info  | oidc.UserInfo | false  | false    | false    | false   | false         | true      | json:"user_info,omitempty"  |          0 |         |
	| is_active  | bool          | false  | false    | false    | true    | false         | false     | json:"is_active,omitempty"  |          0 |         |
	| created_at | time.Time     | false  | false    | false    | true    | false         | false     | json:"created_at,omitempty" |          0 |         |
	+------------+---------------+--------+----------+----------+---------+---------------+-----------+-----------------------------+------------+---------+
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	|     Edge      |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	| organizations | Organization | true    | persons | M2M      | false  | true     |         |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	
