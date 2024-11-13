Datomic is a general purpose database system designed for data-of-record applications

A Datomic database is a set of immutable atomic facts called datoms
[An atomic fact in a database, composed of entity/attribute/value/transaction/added. Pronounced like "datum", but pluralized as datoms.]

Datomic transactions add datoms, never updating or removing them

An atomic unit of work in a database. All Datomic writes are transactional, fully serialized, and ACID (Atomic, Consistent, Isolated, and Durable).