grammar Grammar;

program    : statement* EOF ;

statement   : varDecl      
            | assignStmt   
            | compoundAssignStmt
            | printStmt   
            | ifStmt 
            | whileStmt
            | forStmt
            | blockStmt 
            | postfixStmt
            ;

varDecl     : VAR IDENTIFIER EQUALS expr ;
assignStmt  : IDENTIFIER EQUALS expr ;
compoundAssignStmt : IDENTIFIER op=(PLUSEQUAL | MINUSEQUAL | STAREQUAL | SLASHEQUAL | MODEQUAL | EXPONENTIALEQUAL) expr ;
printStmt   : PRINT expr ;
ifStmt      : IF LPAREN expr RPAREN thenBranch=blockStmt (ELSE elseBranch=blockStmt)? ;
whileStmt   : WHILE LPAREN expr RPAREN body=blockStmt ;
forStmt     : FOR LPAREN init=forInit SEMICOLON cond=expr SEMICOLON post=forPost RPAREN body=blockStmt ;
blockStmt   : LBRACE statement* RBRACE ;

forInit     : varDecl
            | assignStmt
            ;

forPost   : assignStmt
            | postfixStmt
            ;
            
postfixStmt : IDENTIFIER op=(INC | DEC) ;

expr        : (op=NOT | op=BITNOT) expr  # Unary
            | expr op=EXPONENTIAL expr  # Exponential
            | expr op=(STAR | SLASH | MODULO) expr    # MulDivMod
            | expr op=(PLUS | MINUS) expr    # AddSub
            | expr op=(BITLSHIFT | BITRSHIFT) expr  # BitShift
            | expr op=BITAND expr          # BitAnd
            | expr op=BITXOR expr          # BitXor
            | expr op=BITOR expr           # BitOr
            | expr op=(LESS | GREATER | EQUALEQUAL | BANGEQUAL | LESSEQUAL | GREATEREQUAL) expr  # Comparison
            | expr op=AND expr               # And
            | expr op=OR expr                # Or
            | IDENTIFIER                     # Identifier          
            | NUMBER                         # Number     
            | val=(TRUE | FALSE)             # Boolean
            | STRING                         # String
            | LPAREN expr RPAREN             # Parentheses
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
FOR             : 'for' ;
SEMICOLON       : ';' ;
INC             : '++' ;
DEC             : '--' ;
MODULO          : '%' ;
EXPONENTIAL     : '**' ;

PLUSEQUAL       : '+=' ;
MINUSEQUAL      : '-=' ;
STAREQUAL       : '*=' ;
SLASHEQUAL      : '/=' ;
MODEQUAL        : '%=' ;
EXPONENTIALEQUAL : '**=' ;

BITAND          : '&' ;
BITOR           : '|' ;
BITXOR           : '^' ;
BITLSHIFT       : '<<' ;
BITRSHIFT       : '>>' ;
BITNOT          : '~' ;


PRINT           : 'print' ;
TRUE            : 'true' ;
FALSE           : 'false' ;

AND             : 'and' ;
OR              : 'or' ;
NOT             : 'not' ;

IDENTIFIER      : [a-zA-Z_][a-zA-Z0-9_]* ;
NUMBER          : [0-9]+ ;
STRING          : '"' (~["] | '""')* '"' ;
WS              : [ \t\r\n]+ -> skip ;