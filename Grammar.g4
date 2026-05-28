grammar Grammar;

program    : (statement | funcStmt | terminator)* EOF ;

statement   : (varDecl      
            | assignStmt   
            | arrayAssignStmt
            | compoundAssignStmt
            | exprStmt
            | printStmt   
            | ifStmt 
            | whileStmt
            | forStmt
            | blockStmt 
            | postfixStmt
            | returnStmt) terminator?
            ;

returnStmt  : RETURN expr? ;
terminator   : SEMICOLON | NL ;
funcStmt    : FUNC IDENTIFIER LPAREN (IDENTIFIER (COMMA IDENTIFIER)*)? RPAREN blockStmt ;
exprStmt    : expr ;

varDecl     : VAR IDENTIFIER EQUALS expr ;
assignStmt  : IDENTIFIER EQUALS expr ;
arrayAssignStmt : IDENTIFIER LBRACKET expr RBRACKET EQUALS expr ;
compoundAssignStmt : IDENTIFIER op=(PLUSEQUAL | MINUSEQUAL | STAREQUAL | SLASHEQUAL | MODEQUAL | EXPONENTIALEQUAL | BITANDEQUAL | BITOREQUAL | BITXOREQAUL | BITLSHIFTEQUAL | BITRSHIFTEQUAL) expr ;
printStmt   : PRINT expr ;
ifStmt      : IF LPAREN expr RPAREN thenBranch=blockStmt (ELSE elseBranch=blockStmt)? ;
whileStmt   : WHILE LPAREN expr RPAREN body=blockStmt ;
forStmt     : FOR LPAREN init=forInit SEMICOLON cond=expr SEMICOLON post=forPost RPAREN body=blockStmt ;
blockStmt   : LBRACE (statement | funcStmt | terminator)* RBRACE ;

forInit     : varDecl
            | assignStmt
            ;

forPost   : assignStmt
            | postfixStmt
            ;
            
postfixStmt : IDENTIFIER op=(INC | DEC) ;

expr        : expr LBRACKET expr RBRACKET                       # ArrayIndex
            | IDENTIFIER LPAREN (expr (COMMA expr)*)? RPAREN    # FunctionCall
            | op=(NOT | BITNOT | MINUS) expr                       # Unary
            | expr op=EXPONENTIAL expr                          # Exponential
            | expr op=(STAR | SLASH | MODULO) expr              # MulDivMod
            | expr op=(PLUS | MINUS) expr                       # AddSub
            | expr op=(BITLSHIFT | BITRSHIFT) expr              # BitShift
            | expr op=BITAND expr                               # BitAnd
            | expr op=BITXOR expr                               # BitXor
            | expr op=BITOR expr                                # BitOr
            | expr op=(LESS | GREATER | EQUALEQUAL | BANGEQUAL | LESSEQUAL | GREATEREQUAL) expr  # Comparison
            | expr op=AND expr                                  # And
            | expr op=OR expr                                   # Or
            | LBRACKET (expr (COMMA expr)*)? RBRACKET           # ArrayLiteral
            | IDENTIFIER                                        # Identifier          
            | NUMBER                                            # Number     
            | val=(TRUE | FALSE)                                # Boolean
            | STRING                                            # String
            | LPAREN expr RPAREN                                # Parentheses
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
LBRACKET        : '[' ;
RBRACKET        : ']' ;

IF              : 'if' ;
ELSE            : 'else' ;
WHILE           : 'while' ;
FOR             : 'for' ;

SEMICOLON       : ';' ;
COMMA           : ',' ;

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
BITANDEQUAL     : '&=' ;
BITOREQUAL      : '|=' ;
BITXOREQAUL     : '^=' ;
BITLSHIFTEQUAL  : '<<=' ;
BITRSHIFTEQUAL  : '>>=' ;

FUNC            : 'func' ;
RETURN          : 'return' ;
PRINT           : 'print' ;
TRUE            : 'true' ;
FALSE           : 'false' ;

AND             : 'and' ;
OR              : 'or' ;
NOT             : 'not' ;

IDENTIFIER      : [a-zA-Z_][a-zA-Z0-9_]* ;
NUMBER          : [0-9]+ ('.' [0-9]+)? ;
STRING          : '"' (~["] | '""')* '"' ;
NL              : [\r\n]+ ;
WS              : [ \t]+ -> skip ;