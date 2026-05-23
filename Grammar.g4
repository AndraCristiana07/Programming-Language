grammar Grammar;

program    : statement* EOF ;

statement  : varDecl      
           | assignStmt   
           ;

varDecl    : VAR IDENTIFIER EQUALS expr ;
assignStmt : IDENTIFIER EQUALS expr ;

expr       : expr op=(STAR | SLASH) expr    # MulDiv
           | expr op=(PLUS | MINUS) expr    # AddSub
           | expr op=(LESS | GREATER | EQUALEQUAL | BANGEQUAL | LESSEQUAL | GREATEREQUAL) expr  # Comparison
           | IDENTIFIER                     # Identifier          
           | NUMBER                         # Number     
           | val=(TRUE | FALSE)                 # Boolean
           ;


VAR         : 'var' ;
EQUALS      : '=' ;
PLUS        : '+' ;
MINUS       : '-' ;
STAR        : '*' ;
SLASH       : '/' ;
LESS        : '<' ;
GREATER     : '>' ;
EQUALEQUAL  : '==' ;
BANGEQUAL   : '!=' ;
LESSEQUAL    : '<=' ;
GREATEREQUAL : '>=' ;

TRUE        : 'true' ;
FALSE       : 'false' ;

IDENTIFIER  : [a-zA-Z_][a-zA-Z0-9_]* ;
NUMBER      : [0-9]+ ;
WS          : [ \t\r\n]+ -> skip ;