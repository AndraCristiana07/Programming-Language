grammar Grammar;

program    : statement* EOF ;

statement   : varDecl      
            | assignStmt   
            | printStmt   
            | ifStmt 
            | whileStmt
            | blockStmt 
            ;

varDecl     : VAR IDENTIFIER EQUALS expr ;
assignStmt  : IDENTIFIER EQUALS expr ;
printStmt   : PRINT expr ;
ifStmt      : IF LPAREN expr RPAREN thenBranch=blockStmt (ELSE elseBranch=blockStmt)? ;
whileStmt   : WHILE LPAREN expr RPAREN body=blockStmt ;
blockStmt   : LBRACE statement* RBRACE ;

expr        : expr op=(STAR | SLASH) expr    # MulDiv
            | expr op=(PLUS | MINUS) expr    # AddSub
            | expr op=(LESS | GREATER | EQUALEQUAL | BANGEQUAL | LESSEQUAL | GREATEREQUAL) expr  # Comparison
            | IDENTIFIER                     # Identifier          
            | NUMBER                         # Number     
            | val=(TRUE | FALSE)             # Boolean
            | STRING                         # String
            ;


VAR             : 'var' ;
EQUALS          : '=' ;
PLUS            : '+' ;
MINUS           : '-' ;
STAR            : '*' ;
SLASH           : '/' ;
LESS            : '<' ;
GREATER         : '>' ;
EQUALEQUAL      : '==' ;
BANGEQUAL       : '!=' ;
LESSEQUAL       : '<=' ;
GREATEREQUAL    : '>=' ;
LBRACE          : '{' ;
LPAREN          : '(' ;
RPAREN          : ')' ;
RBRACE          : '}' ;
IF              : 'if' ;
ELSE            : 'else' ;
WHILE           : 'while' ;

PRINT           : 'print' ;
TRUE            : 'true' ;
FALSE           : 'false' ;

IDENTIFIER      : [a-zA-Z_][a-zA-Z0-9_]* ;
NUMBER          : [0-9]+ ;
STRING          : '"' (~["] | '""')* '"' ;
WS              : [ \t\r\n]+ -> skip ;