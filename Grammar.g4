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
            | switchStmt
            | forInStmt
            | forStmt
            | blockStmt 
            | postfixStmt
            | tryCatchStmt
            | throwStmt
            | returnStmt
            | breakStmt        
            | continueStmt) terminator?
            ;

returnStmt          : RETURN expr? ;
terminator          : SEMICOLON | NL ;
funcStmt            : FUNC IDENTIFIER LPAREN (IDENTIFIER (COMMA IDENTIFIER)*)? RPAREN blockStmt ;
exprStmt            : expr ;
tryCatchStmt        : TRY tryBody=blockStmt CATCH LPAREN IDENTIFIER RPAREN catchBody=blockStmt ;
throwStmt           : THROW expr ;
breakStmt           : BREAK;
continueStmt        : CONTINUE;
switchStmt          : SWITCH LPAREN expr RPAREN LBRACE terminator* caseBlock* defaultBlock?terminator* RBRACE ;

caseBlock           : CASE expr COLON (statement | terminator)* ;
defaultBlock        : DEFAULT COLON (statement | terminator)* ;

forInStmt           : FOR LPAREN VAR? id=IDENTIFIER IN expr RPAREN body=blockStmt ;

varDecl             : VAR IDENTIFIER EQUALS expr ;
assignStmt          : expr EQUALS expr ;
arrayAssignStmt     : IDENTIFIER LBRACKET expr RBRACKET EQUALS expr ;
compoundAssignStmt  : expr op=(PLUSEQUAL | MINUSEQUAL | STAREQUAL | SLASHEQUAL | MODEQUAL | EXPONENTIALEQUAL | BITANDEQUAL | BITOREQUAL | BITXOREQAUL | BITLSHIFTEQUAL | BITRSHIFTEQUAL) expr ;
printStmt           : PRINT expr ;
ifStmt              : IF LPAREN expr RPAREN thenBranch=blockStmt (ELIF LPAREN elifCond+=expr RPAREN elifBranch+=blockStmt)* (ELSE elseBranch=blockStmt)? ;
whileStmt           : WHILE LPAREN expr RPAREN body=blockStmt ;
forStmt             : FOR LPAREN init=forInit SEMICOLON cond=expr SEMICOLON post=forPost RPAREN body=blockStmt ;
blockStmt           : LBRACE (statement | funcStmt | terminator)* RBRACE ;

forInit     : varDecl
            | assignStmt
            ;

forPost   : assignStmt
            | postfixStmt
            ;
            
postfixStmt        : expr op=(INC | DEC) ;

expr        : expr LBRACKET expr RBRACKET                           # ArrayIndex
            | expr DOT IDENTIFIER                                   # FieldAccess
            | IDENTIFIER LPAREN (expr (COMMA expr)*)? RPAREN        # FunctionCall
            | op=(NOT | BITNOT | MINUS) expr                        # Unary
            | expr op=EXPONENTIAL expr                              # Exponential
            | expr op=(STAR | SLASH | MODULO) expr                  # MulDivMod
            | expr op=(PLUS | MINUS) expr                           # AddSub
            | expr op=(BITLSHIFT | BITRSHIFT) expr                  # BitShift
            | expr op=BITAND expr                                   # BitAnd
            | expr op=BITXOR expr                                   # BitXor
            | expr op=BITOR expr                                    # BitOr
            | expr op=(LESS | GREATER | EQUALEQUAL | BANGEQUAL | LESSEQUAL | GREATEREQUAL) expr  # Comparison
            | expr op=AND expr                                      # And
            | expr op=OR expr                                       # Or
            | trueExpr=expr IF condExpr=expr ELSE falseExpr=expr    # TernaryOp
            | LBRACKET (expr (COMMA expr)*)? RBRACKET               # ArrayLiteral
            | LBRACE (mapEntry (COMMA mapEntry)*)? RBRACE           # MapLiteral                              
            | IDENTIFIER                                            # Identifier          
            | NUMBER                                                # Number     
            | val=(TRUE | FALSE)                                    # Boolean
            | STRING                                                # String
            | 'null'                                                # Null
            | LPAREN expr RPAREN                                    # Parentheses
            ;

mapEntry : expr COLON expr ;

SWITCH          : 'switch' ;
CASE            : 'case' ;
DEFAULT         : 'default' ;
TRY             : 'try' ;
CATCH           : 'catch' ;
THROW           : 'throw' ;
IN              : 'in' ;
DOT             : '.' ;
COLON           : ':' ;
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
ELIF            : 'elif' ;
WHILE           : 'while' ;
FOR             : 'for' ;
BREAK           : 'break' ;
CONTINUE        : 'continue' ;

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
LINE_COMMENT
    : '//' ~[\r\n]* -> skip
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> skip
    ;
NL              : [\r\n]+ ;
WS              : [ \t]+ -> skip ;
