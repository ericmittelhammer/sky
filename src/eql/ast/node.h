#ifndef _eql_ast_node_h
#define _eql_ast_node_h

#include "int_literal.h"
#include "float_literal.h"

//==============================================================================
//
// Definitions
//
//==============================================================================

// Defines the types of expressions available.
typedef enum {
    EQL_AST_TYPE_INT_LITERAL,
    EQL_AST_TYPE_FLOAT_LITERAL
} eql_ast_node_type_e;

// Represents an node in the AST.
typedef struct eql_ast_node {
    eql_ast_node_type_e type;
    union {
        eql_ast_int_literal int_literal;
        eql_ast_float_literal float_literal;
    };
} eql_ast_node;


//==============================================================================
//
// Functions
//
//==============================================================================

void eql_ast_node_free(eql_ast_node *node);

#endif