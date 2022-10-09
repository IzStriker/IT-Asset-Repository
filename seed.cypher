// Document
CREATE(document:AssetType{name:"Document"})
CREATE(extension:AssetAttributeType{name:"Extension"})
CREATE(size:AssetAttributeType{name:"Size"})
CREATE(filename:AssetAttributeType{name:"Filename"})
CREATE(location:AssetAttributeType{name:"Location"})
CREATE 
    (document)-[:OWNS]->(extension),
    (document)-[:OWNS]->(size),
    (document)-[:OWNS]->(filename),
    (document)-[:OWNS]->(location)

// Code Document 
CREATE(codeDocument:AssetType{name:"Code Document"})
CREATE(loc:AssetAttributeType{name:"Line of Code"})
CREATE(language:AssetAttributeType{name:"Programming Language"})
CREATE
    (codeDocument)-[:EXTENDS]->(document),
    (codeDocument)-[:OWNS]->(loc),
    (codeDocument)-[:OWNS]->(language)

// Configuration Document 
CREATE(configDocument:AssetType{name:"Config Document"})
CREATE(type:AssetAttributeType{name:"Type"})
CREATE
    (configDocument)-[:EXTENDS]->(document),
    (configDocument)-[:OWNS]->(type)

// Java Class
CREATE(class:AssetType{name:"Java Class"})
CREATE(className:AssetAttributeType{name:"Class Name"})
CREATE 
    (class)-[:EXTENDS]->(codeDocument),
    (class)-[:OWNS]->(className)