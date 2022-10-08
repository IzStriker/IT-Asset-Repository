// Document
CREATE(document:ASSETTYPE{name:"Document"})
CREATE(extension:ATTRIBUTE{name:"Extension"})
CREATE(size:ATTRIBUTE{name:"Size"})
CREATE(filename:ATTRIBUTE{name:"Filename"})
CREATE(location:ATTRIBUTE{name:"Location"})
CREATE 
    (document)-[:OWNS]->(extension),
    (document)-[:OWNS]->(size),
    (document)-[:OWNS]->(filename),
    (document)-[:OWNS]->(location)

// Code Document 
CREATE(codeDocument:ASSETTYPE{name:"Code Document"})
CREATE(loc:ATTRIBUTE{name:"Line of Code"})
CREATE(language:ATTRIBUTE{name:"Programming Language"})
CREATE
    (codeDocument)-[:EXTENDS]->(document),
    (codeDocument)-[:OWNS]->(loc),
    (codeDocument)-[:language]->(language)

// Configuration Document 
CREATE(configDocument:ASSETTYPE{name:"Config Document"})
CREATE(type:ATTRIBUTE{name:"Type"})
CREATE
    (configDocument)-[:EXTENDS]->(document),
    (configDocument)-[:OWNS]->(type)

// Java Class
CREATE(class:ASSETTYPE{name:"Java Class"})
CREATE(className:ATTRIBUTE{name:"Class Name"})
CREATE 
    (class)-[:EXTENDS]->(codeDocument),
    (class)-[:OWNS]->(className)