<!--
ignore these words in spell check for this file
// cSpell:ignore Silberschatz Korth Sudarshan ntile cume_dist arity ODBC JDBC OLAP sqsuspet sqsubpet sqsupbset
-->

# Database System Concepts

<!-- <details> -->
<summary>
Database System Concepts book.
</summary>

I'm reading a god-damn book about SQL?

Authors: Abraham Silberschatz, Henry F. Korth, S. Sudarshan.

## Chapter 3 - Introduction To SQL

- relations
- basic sql query structure
  - `select`
  - `from`
  - `where`
  - `order`
- Natural Join
  - `as`
- Set operations
  - `union`
  - `intersect`
  - `except`
  - set comparisons
- Null value (truthiness)
- Aggregate functions
  - `avg`, `min`, `max`, `sum`, `count`
  - `group by`
  - `having`
- sub queries

## Chapter 4 - Intermediate SQL

- Joins
- Views
  - Materialized views
- Transactions
  - `Commit`
  - `Rollback`
- Constraints
- `Check` Clause (arbitrary predicate)
- Referential Integrity
- Data Types
  - user defined types
  - domains
- Default values
- Catalogs and Schemas
- Authorization
  - privileges
  - roles

## Chapter 5 - Advanced SQL

- Accessing SQL from software
  - Dynamic SQL - queries are strings and are processed by the database
  - Embedded SQL - queries are compiled by the client
  - Prepared Statements - faster and safer
  - Callable Statements
- Metadata Features
- ODBC and JDBC
- SQL functions and procedures
- Exception conditions and handlers
- Triggers
  - event, condition, action
  - before or after event
  - per row or per table
- Recursive Queries
  - transitive closures
- Advanced aggregations
  - `rank`, `dense_rank`,
  - `percent_rank`, `ntile`
  - `cume_dist`
  - `row_number`
  - `over`
  - `partition by`
- Windowing
  - `preceding`, `following`
  - `unbounded`
  - `current row`
  - `range between`
- OLAP - online analytical processing - analytics solutions and tools
  - `pivot`(cross tabs)
  - `cube`, `rollup` (generalizations of `group by`)
  - `decode`

## Chapter 6 - Formal Relational Query Languages

- operations - can act on one relation and produce one relation back (unary), or run on pairs of relations (binary)
  - select (unary) - select/filter tuples that satisfy a predicate
  - project (unary) - project relation into new form with some of the attributes. will remove duplicates.
  - union (binary) - combine two relations together. will remove duplicates. relations must be the same arity (number of attributes) and the attributes must be the same domain.
  - set difference (binary) - tuples in one relation but not another. order matters.
  - Cartesian product (binary) - create a new relation with a schema that is the union of the two relations schema. each tuple gets paired with all the tuples in the other relations.
  - rename (unary) - give temporary name to a relation as part of the formula (????)
  - set intersection - appearing in both
  - natural join - Cartesian product which is filtered on some natural common attribute
  - assignment - give temporary name (???)
  - outer join - dealing with missing information (right outer join, left outer join, full outer join)
- aggregations
- MultiSets (may contain duplicates)
- Tuple relational calculus - another type of syntax/form to write algebra
- Domain relation calculus.

the format for unary operator is the symbol, and then the predicate in subscript and finally the relation name in parentheses.

[symbols](https://www.cs.uleth.ca/~rice/latex/worksheet.pdf)

| Symbol name | LATEXcode |symbol|
|---|---|---|
| leftarrow | `\leftarrow` | $\leftarrow$ |
| select | `\sigma` | $\sigma$ |
| project | `\Pi` | $\Pi$ |
| inner join | `\bowtie` | $\bowtie$ |
| left outer join | `\sqsubpet\bowtie` | $\sqsupset\bowtie$ |
| right outer join | `\bowtie\sqsuspet` | $\bowtie\sqsubset$ |
| full outer join | `\sqsupbset\bowtie\sqsuspet` | $\sqsupset\bowtie\sqsubset$ |
| cross product | `\times` | $\times$ |
| rename | `\rho` | $\rho$ |
| less than | `<` | $<$ |
| greater than |`>` | $>$ |
| less than or equal | `\leq`  | $\leq$ |
| greater than or equal | `\geq` | $\geq$ |
| equal | `=`  | $=$ |
| not equal | `\neq` | $\neq$ |
| and | `\wedge` | $\wedge$ |
| or | `\vee` | $\vee$ |
| not | `\neg` | $\neg$ |

$\sigma_{department\_name="Physics"}(instructor)$
$\Pi_{ID, Name, Salary}(instructor)$

combining together to find the names of all instructors from the physics department. since the result of all operations on a relation is a relation, we can use them inside the parentheses.

$\Pi_{name}(\sigma_{department\_name="Physics"}(instructor))$

the `set intersection`, `natural join` and `assignment` operators don't add power over the fundamental operators, but make our life easier.

Extended relation algebra operations:

have the format of symbol, list of arguments which are attributes or constants, and then the relations.

we use them for aggregation, with the calligraphic G denoting aggregation `\mathcal{G}` $\mathcal{G}$. the name of the aggregation is written as it (sum, max, min, count, etc...). we can denote grouping by prefixing the symbol with the attribute we want to group on.

## Chapter 7 - Database Design and the E-R Model

E-R -> Entity Relationship, focusing on how to choose a schema. which entities are present, what are the main attributes and what relations they have to each other. specifying the required operations on the data.\
entities are the "things" the system deals with - people, places, products, etc...

_Redundancy_ is bad for schema design. data should only appear once, and should not be replicated across entities. if it is replicated, the different representations of the data may become inconsistent across time. _Incompleteness_ is another product of bad design. we might end up creating workaround items that belong to one entity and have missing data, because the entity actually covers more than one distinct "thing".

the entity-relationship model has three concepts:

- entity sets
- relationship sets
- attributes

if an entity is a distinct "thing", the entity set is a collection of those different things together, while they still belong to same category of "things". entities have attributes, so while all the entities in set have the same attributes, the value of those attributes differs.

a relationship is an association between entities.
 the entities are "participants" in the relationship, and we can define them as having a "role", relationships can also have attributes.

attributes themselves have domains (value sets), or possible allowed values. this can numbers, dates, text, etc.\
we can have simple attributes such as name or address, or composite attributes - first and last name, street, city, state and zip code. in some cases this makes sense, and in some cases it doesn't. (it makes sense to query for all students from a certain state, but less so to query for all student with the same first name). attributes can also be _single-valued_ or _multi-valued_. an id is single value, any entity can only have one. but a phone number is a multi-valued attribute, since a single person can have multiple numbers. a derived attribute is something that we can infer from from other relationships or entities, such as deriving someone's age from the date of birth. we don't store the derived values, we compute them instead.

a null value can mean the data is missing for the attribute (unknown) or doesn't exist at all.

schemas have constrains, such as relationships mapping cardinality - how many times can an entity be in the role in the relationship set.

- one to one
- one to many
- many to one
- many to many

participation can be _total_ if every entity from the entity set participates in a relationship (every person has at least one person as a parent), or _partial_ if some entities don't participate in it (not every person owns a car).

entities are different from one another because of their "keys", special attributes that must uniquely identify the entity from all other entities. the key can be one or more attributes, as long as the set of the attributes is unique. this is also true for relationships, they also have some unique way to identify them from one another.

- super key
- candidate key
- primary key

when we define the schema, we remove all the attributes which aren't unique, that means that we also remove attributes which refer to other entities, since those are conceptually relations, rather than true attributes.

### Diagram Design

> - Rectangles for entity sets, the name of the entity set and the attributes, attributes which are part of the primary key are underlined.
> - Diamonds represent relationship sets.
> - Lines link between entity sets and relationship sets.
>   - dashed lines link attributes of the relationship sets to the relationship set.
>   - double lines indicate total participation of an an entity in a relationship set (entity must appear in the relationship set)
> - Double Diamond represent identifying relationship sets and weak entity sets.

When denoting entity attributes, we can use indentation to mark attributes as complex - meaning that are composed of internal values, for example, the address is composed of street, city, state, and zip code attributes. we use `{attribute_name}` to denote that the attribute can have more than one value (like phone number), and `attribute_name()` to denote a derived attribute (such as age, which is calculated from data of birth).

the linking lines can also show cardinality (mapping).

- one-to-one - both sides of the line have an arrow head.
- one-to-many or many-to-one - the side with the single entity has the arrow head.
- many-to-many - no arrow head on either side of the line.

alternatively, we can write complex cardinality rules by denoting number on the lines. we write the minimum and maximum number of times each entity appears in the relationship set, with `*` as the symbol for unlimited. having a minimum value of 1 means that the entity must be part of the relationship set, at least one time. a minimal value of zero means that the entity doesn't necessary appear in the relationship set.

- `1..1` - must appear once, no more, no less.
- `0..1` - can appear at most once, but doesn't have to appear
- `1..*` - must appear, can appear any number of times

Links between entity sets and relationship sets can also show roles as text over the links, this is especially useful when a relationship sets  links between entities of the same type. Relationship sets don't have to be binary, they can involve more than two entity sets, but that makes marking cardinality more difficult.

Weak Entity sets are entity sets which don't have enough attributes to define a primary key, and are not meaningful enough on their own. instead, they have partial keys as discriminators, which are marked with a dash underline. the weak entity only has meaning through the identifying relationship set (double diamond marking).

---
we can either use diagrams or entity-relation schemas to model our objects, and there are decisions which we must take whether a thing is an entity with a relationship set, or simply an attribute of another entity set. some relations can be modeled as entities on their own.

### Extended Entity-Relation Features

some databases allow for additional ways of modeling, such as specialization and generalization. specialization happens when designating some entities in an entity set as similar to one another, and as having additional attributes unique only to them. they still have all the same attributes as other entities, but are considered distinct in some use-cases. this is equivalent to OOP **Inheritance**, a "Is-a" relationship, denoted with an hollow (empty) arrow head between the rectangles of the diagram (UML!). in some cases, an entity can belong to more than one specialized entity sets (multiple inheritance).\
Generalization is the process of taking multiple distinct entitySets and grouping them together because they all share some common relevant attributes, this happen "bottom-up".\
Assigning an entity from a high-level entitySet to a lower (specialized) set can be done based on attributes or directly (case-by-case basis), an entity might belong to more than one lower-level specializations, and it might be required that it belongs to one, and only one.

Aggregation is the process of treating a relationship set as a higher level entity.

## Chapter 8 - Relational Database Design

</details>
