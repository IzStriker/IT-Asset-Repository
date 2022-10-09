// Document
CREATE(document:AssetType{name:"Document"})
CREATE(extension:AssetAttributeType{name:"Extension", type: "STRING"})
CREATE(size:AssetAttributeType{name:"Size", type: "NUMBER"})
CREATE(filename:AssetAttributeType{name:"Filename", type: "STRING"})
CREATE(location:AssetAttributeType{name:"Location", type: "STRING"})
CREATE 
    (document)-[:OWNS]->(extension),
    (document)-[:OWNS]->(size),
    (document)-[:OWNS]->(filename),
    (document)-[:OWNS]->(location)

// Code Document 
CREATE(codeDocument:AssetType{name:"Code Document"})
CREATE(loc:AssetAttributeType{name:"Line of Code", type: "NUMBER"})
CREATE(language:AssetAttributeType{name:"Programming Language", type: "STRING"})
CREATE
    (codeDocument)-[:EXTENDS]->(document),
    (codeDocument)-[:OWNS]->(loc),
    (codeDocument)-[:OWNS]->(language)

// Configuration Document 
CREATE(configDocument:AssetType{name:"Config Document"})
CREATE(type:AssetAttributeType{name:"Type", type: "STRING"})
CREATE
    (configDocument)-[:EXTENDS]->(document),
    (configDocument)-[:OWNS]->(type)

// Java Class
CREATE(class:AssetType{name:"Java Class"})
CREATE(className:AssetAttributeType{name:"Class Name", type: "STRING"})
CREATE 
    (class)-[:EXTENDS]->(codeDocument),
    (class)-[:OWNS]->(className)