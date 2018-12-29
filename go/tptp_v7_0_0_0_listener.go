// Code generated from tptp_v7_0_0_0.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tptp_v7_0_0_0

import "github.com/antlr/antlr4/runtime/Go/antlr"

// tptp_v7_0_0_0Listener is a complete listener for a parse tree produced by tptp_v7_0_0_0Parser.
type tptp_v7_0_0_0Listener interface {
	antlr.ParseTreeListener

	// EnterTptp_file is called when entering the tptp_file production.
	EnterTptp_file(c *Tptp_fileContext)

	// EnterTptp_input is called when entering the tptp_input production.
	EnterTptp_input(c *Tptp_inputContext)

	// EnterAnnotated_formula is called when entering the annotated_formula production.
	EnterAnnotated_formula(c *Annotated_formulaContext)

	// EnterTpi_annotated is called when entering the tpi_annotated production.
	EnterTpi_annotated(c *Tpi_annotatedContext)

	// EnterTpi_formula is called when entering the tpi_formula production.
	EnterTpi_formula(c *Tpi_formulaContext)

	// EnterThf_annotated is called when entering the thf_annotated production.
	EnterThf_annotated(c *Thf_annotatedContext)

	// EnterTfx_annotated is called when entering the tfx_annotated production.
	EnterTfx_annotated(c *Tfx_annotatedContext)

	// EnterTff_annotated is called when entering the tff_annotated production.
	EnterTff_annotated(c *Tff_annotatedContext)

	// EnterTcf_annotated is called when entering the tcf_annotated production.
	EnterTcf_annotated(c *Tcf_annotatedContext)

	// EnterFof_annotated is called when entering the fof_annotated production.
	EnterFof_annotated(c *Fof_annotatedContext)

	// EnterCnf_annotated is called when entering the cnf_annotated production.
	EnterCnf_annotated(c *Cnf_annotatedContext)

	// EnterAnnotations is called when entering the annotations production.
	EnterAnnotations(c *AnnotationsContext)

	// EnterFormula_role is called when entering the formula_role production.
	EnterFormula_role(c *Formula_roleContext)

	// EnterThf_formula is called when entering the thf_formula production.
	EnterThf_formula(c *Thf_formulaContext)

	// EnterThf_logic_formula is called when entering the thf_logic_formula production.
	EnterThf_logic_formula(c *Thf_logic_formulaContext)

	// EnterThf_binary_formula is called when entering the thf_binary_formula production.
	EnterThf_binary_formula(c *Thf_binary_formulaContext)

	// EnterThf_binary_pair is called when entering the thf_binary_pair production.
	EnterThf_binary_pair(c *Thf_binary_pairContext)

	// EnterThf_binary_tuple is called when entering the thf_binary_tuple production.
	EnterThf_binary_tuple(c *Thf_binary_tupleContext)

	// EnterThf_or_formula is called when entering the thf_or_formula production.
	EnterThf_or_formula(c *Thf_or_formulaContext)

	// EnterThf_and_formula is called when entering the thf_and_formula production.
	EnterThf_and_formula(c *Thf_and_formulaContext)

	// EnterThf_apply_formula is called when entering the thf_apply_formula production.
	EnterThf_apply_formula(c *Thf_apply_formulaContext)

	// EnterThf_unitary_formula is called when entering the thf_unitary_formula production.
	EnterThf_unitary_formula(c *Thf_unitary_formulaContext)

	// EnterThf_quantified_formula is called when entering the thf_quantified_formula production.
	EnterThf_quantified_formula(c *Thf_quantified_formulaContext)

	// EnterThf_quantification is called when entering the thf_quantification production.
	EnterThf_quantification(c *Thf_quantificationContext)

	// EnterThf_variable_list is called when entering the thf_variable_list production.
	EnterThf_variable_list(c *Thf_variable_listContext)

	// EnterThf_variable is called when entering the thf_variable production.
	EnterThf_variable(c *Thf_variableContext)

	// EnterThf_typed_variable is called when entering the thf_typed_variable production.
	EnterThf_typed_variable(c *Thf_typed_variableContext)

	// EnterThf_unary_formula is called when entering the thf_unary_formula production.
	EnterThf_unary_formula(c *Thf_unary_formulaContext)

	// EnterThf_atom is called when entering the thf_atom production.
	EnterThf_atom(c *Thf_atomContext)

	// EnterThf_function is called when entering the thf_function production.
	EnterThf_function(c *Thf_functionContext)

	// EnterThf_conn_term is called when entering the thf_conn_term production.
	EnterThf_conn_term(c *Thf_conn_termContext)

	// EnterThf_conditional is called when entering the thf_conditional production.
	EnterThf_conditional(c *Thf_conditionalContext)

	// EnterThf_let is called when entering the thf_let production.
	EnterThf_let(c *Thf_letContext)

	// EnterThf_arguments is called when entering the thf_arguments production.
	EnterThf_arguments(c *Thf_argumentsContext)

	// EnterThf_type_formula is called when entering the thf_type_formula production.
	EnterThf_type_formula(c *Thf_type_formulaContext)

	// EnterThf_typeable_formula is called when entering the thf_typeable_formula production.
	EnterThf_typeable_formula(c *Thf_typeable_formulaContext)

	// EnterThf_subtype is called when entering the thf_subtype production.
	EnterThf_subtype(c *Thf_subtypeContext)

	// EnterThf_top_level_type is called when entering the thf_top_level_type production.
	EnterThf_top_level_type(c *Thf_top_level_typeContext)

	// EnterThf_unitary_type is called when entering the thf_unitary_type production.
	EnterThf_unitary_type(c *Thf_unitary_typeContext)

	// EnterThf_apply_type is called when entering the thf_apply_type production.
	EnterThf_apply_type(c *Thf_apply_typeContext)

	// EnterThf_binary_type is called when entering the thf_binary_type production.
	EnterThf_binary_type(c *Thf_binary_typeContext)

	// EnterThf_mapping_type is called when entering the thf_mapping_type production.
	EnterThf_mapping_type(c *Thf_mapping_typeContext)

	// EnterThf_xprod_type is called when entering the thf_xprod_type production.
	EnterThf_xprod_type(c *Thf_xprod_typeContext)

	// EnterThf_union_type is called when entering the thf_union_type production.
	EnterThf_union_type(c *Thf_union_typeContext)

	// EnterThf_sequent is called when entering the thf_sequent production.
	EnterThf_sequent(c *Thf_sequentContext)

	// EnterThf_tuple is called when entering the thf_tuple production.
	EnterThf_tuple(c *Thf_tupleContext)

	// EnterThf_formula_list is called when entering the thf_formula_list production.
	EnterThf_formula_list(c *Thf_formula_listContext)

	// EnterTfx_formula is called when entering the tfx_formula production.
	EnterTfx_formula(c *Tfx_formulaContext)

	// EnterTfx_logic_formula is called when entering the tfx_logic_formula production.
	EnterTfx_logic_formula(c *Tfx_logic_formulaContext)

	// EnterTff_formula is called when entering the tff_formula production.
	EnterTff_formula(c *Tff_formulaContext)

	// EnterTff_logic_formula is called when entering the tff_logic_formula production.
	EnterTff_logic_formula(c *Tff_logic_formulaContext)

	// EnterTff_binary_formula is called when entering the tff_binary_formula production.
	EnterTff_binary_formula(c *Tff_binary_formulaContext)

	// EnterTff_binary_nonassoc is called when entering the tff_binary_nonassoc production.
	EnterTff_binary_nonassoc(c *Tff_binary_nonassocContext)

	// EnterTff_binary_assoc is called when entering the tff_binary_assoc production.
	EnterTff_binary_assoc(c *Tff_binary_assocContext)

	// EnterTff_or_formula is called when entering the tff_or_formula production.
	EnterTff_or_formula(c *Tff_or_formulaContext)

	// EnterTff_and_formula is called when entering the tff_and_formula production.
	EnterTff_and_formula(c *Tff_and_formulaContext)

	// EnterTff_unitary_formula is called when entering the tff_unitary_formula production.
	EnterTff_unitary_formula(c *Tff_unitary_formulaContext)

	// EnterTff_quantified_formula is called when entering the tff_quantified_formula production.
	EnterTff_quantified_formula(c *Tff_quantified_formulaContext)

	// EnterTff_variable_list is called when entering the tff_variable_list production.
	EnterTff_variable_list(c *Tff_variable_listContext)

	// EnterTff_variable is called when entering the tff_variable production.
	EnterTff_variable(c *Tff_variableContext)

	// EnterTff_typed_variable is called when entering the tff_typed_variable production.
	EnterTff_typed_variable(c *Tff_typed_variableContext)

	// EnterTff_unary_formula is called when entering the tff_unary_formula production.
	EnterTff_unary_formula(c *Tff_unary_formulaContext)

	// EnterTff_atomic_formula is called when entering the tff_atomic_formula production.
	EnterTff_atomic_formula(c *Tff_atomic_formulaContext)

	// EnterTff_conditional is called when entering the tff_conditional production.
	EnterTff_conditional(c *Tff_conditionalContext)

	// EnterTff_let is called when entering the tff_let production.
	EnterTff_let(c *Tff_letContext)

	// EnterTff_let_term_defns is called when entering the tff_let_term_defns production.
	EnterTff_let_term_defns(c *Tff_let_term_defnsContext)

	// EnterTff_let_term_list is called when entering the tff_let_term_list production.
	EnterTff_let_term_list(c *Tff_let_term_listContext)

	// EnterTff_let_term_defn is called when entering the tff_let_term_defn production.
	EnterTff_let_term_defn(c *Tff_let_term_defnContext)

	// EnterTff_let_term_binding is called when entering the tff_let_term_binding production.
	EnterTff_let_term_binding(c *Tff_let_term_bindingContext)

	// EnterTff_let_formula_defns is called when entering the tff_let_formula_defns production.
	EnterTff_let_formula_defns(c *Tff_let_formula_defnsContext)

	// EnterTff_let_formula_list is called when entering the tff_let_formula_list production.
	EnterTff_let_formula_list(c *Tff_let_formula_listContext)

	// EnterTff_let_formula_defn is called when entering the tff_let_formula_defn production.
	EnterTff_let_formula_defn(c *Tff_let_formula_defnContext)

	// EnterTff_let_formula_binding is called when entering the tff_let_formula_binding production.
	EnterTff_let_formula_binding(c *Tff_let_formula_bindingContext)

	// EnterTff_sequent is called when entering the tff_sequent production.
	EnterTff_sequent(c *Tff_sequentContext)

	// EnterTff_formula_tuple is called when entering the tff_formula_tuple production.
	EnterTff_formula_tuple(c *Tff_formula_tupleContext)

	// EnterTff_formula_tuple_list is called when entering the tff_formula_tuple_list production.
	EnterTff_formula_tuple_list(c *Tff_formula_tuple_listContext)

	// EnterTff_typed_atom is called when entering the tff_typed_atom production.
	EnterTff_typed_atom(c *Tff_typed_atomContext)

	// EnterTff_subtype is called when entering the tff_subtype production.
	EnterTff_subtype(c *Tff_subtypeContext)

	// EnterTff_top_level_type is called when entering the tff_top_level_type production.
	EnterTff_top_level_type(c *Tff_top_level_typeContext)

	// EnterTf1_quantified_type is called when entering the tf1_quantified_type production.
	EnterTf1_quantified_type(c *Tf1_quantified_typeContext)

	// EnterTff_monotype is called when entering the tff_monotype production.
	EnterTff_monotype(c *Tff_monotypeContext)

	// EnterTff_unitary_type is called when entering the tff_unitary_type production.
	EnterTff_unitary_type(c *Tff_unitary_typeContext)

	// EnterTff_atomic_type is called when entering the tff_atomic_type production.
	EnterTff_atomic_type(c *Tff_atomic_typeContext)

	// EnterTff_type_arguments is called when entering the tff_type_arguments production.
	EnterTff_type_arguments(c *Tff_type_argumentsContext)

	// EnterTff_mapping_type is called when entering the tff_mapping_type production.
	EnterTff_mapping_type(c *Tff_mapping_typeContext)

	// EnterTff_xprod_type is called when entering the tff_xprod_type production.
	EnterTff_xprod_type(c *Tff_xprod_typeContext)

	// EnterTcf_formula is called when entering the tcf_formula production.
	EnterTcf_formula(c *Tcf_formulaContext)

	// EnterTcf_logic_formula is called when entering the tcf_logic_formula production.
	EnterTcf_logic_formula(c *Tcf_logic_formulaContext)

	// EnterTcf_quantified_formula is called when entering the tcf_quantified_formula production.
	EnterTcf_quantified_formula(c *Tcf_quantified_formulaContext)

	// EnterFof_formula is called when entering the fof_formula production.
	EnterFof_formula(c *Fof_formulaContext)

	// EnterFof_logic_formula is called when entering the fof_logic_formula production.
	EnterFof_logic_formula(c *Fof_logic_formulaContext)

	// EnterFof_binary_formula is called when entering the fof_binary_formula production.
	EnterFof_binary_formula(c *Fof_binary_formulaContext)

	// EnterFof_binary_nonassoc is called when entering the fof_binary_nonassoc production.
	EnterFof_binary_nonassoc(c *Fof_binary_nonassocContext)

	// EnterFof_binary_assoc is called when entering the fof_binary_assoc production.
	EnterFof_binary_assoc(c *Fof_binary_assocContext)

	// EnterFof_or_formula is called when entering the fof_or_formula production.
	EnterFof_or_formula(c *Fof_or_formulaContext)

	// EnterFof_and_formula is called when entering the fof_and_formula production.
	EnterFof_and_formula(c *Fof_and_formulaContext)

	// EnterFof_unitary_formula is called when entering the fof_unitary_formula production.
	EnterFof_unitary_formula(c *Fof_unitary_formulaContext)

	// EnterFof_quantified_formula is called when entering the fof_quantified_formula production.
	EnterFof_quantified_formula(c *Fof_quantified_formulaContext)

	// EnterFof_variable_list is called when entering the fof_variable_list production.
	EnterFof_variable_list(c *Fof_variable_listContext)

	// EnterFof_unary_formula is called when entering the fof_unary_formula production.
	EnterFof_unary_formula(c *Fof_unary_formulaContext)

	// EnterFof_infix_unary is called when entering the fof_infix_unary production.
	EnterFof_infix_unary(c *Fof_infix_unaryContext)

	// EnterFof_atomic_formula is called when entering the fof_atomic_formula production.
	EnterFof_atomic_formula(c *Fof_atomic_formulaContext)

	// EnterFof_plain_atomic_formula is called when entering the fof_plain_atomic_formula production.
	EnterFof_plain_atomic_formula(c *Fof_plain_atomic_formulaContext)

	// EnterFof_defined_atomic_formula is called when entering the fof_defined_atomic_formula production.
	EnterFof_defined_atomic_formula(c *Fof_defined_atomic_formulaContext)

	// EnterFof_defined_plain_formula is called when entering the fof_defined_plain_formula production.
	EnterFof_defined_plain_formula(c *Fof_defined_plain_formulaContext)

	// EnterFof_defined_infix_formula is called when entering the fof_defined_infix_formula production.
	EnterFof_defined_infix_formula(c *Fof_defined_infix_formulaContext)

	// EnterFof_system_atomic_formula is called when entering the fof_system_atomic_formula production.
	EnterFof_system_atomic_formula(c *Fof_system_atomic_formulaContext)

	// EnterFof_plain_term is called when entering the fof_plain_term production.
	EnterFof_plain_term(c *Fof_plain_termContext)

	// EnterFof_defined_term is called when entering the fof_defined_term production.
	EnterFof_defined_term(c *Fof_defined_termContext)

	// EnterFof_defined_atomic_term is called when entering the fof_defined_atomic_term production.
	EnterFof_defined_atomic_term(c *Fof_defined_atomic_termContext)

	// EnterFof_defined_plain_term is called when entering the fof_defined_plain_term production.
	EnterFof_defined_plain_term(c *Fof_defined_plain_termContext)

	// EnterFof_system_term is called when entering the fof_system_term production.
	EnterFof_system_term(c *Fof_system_termContext)

	// EnterFof_arguments is called when entering the fof_arguments production.
	EnterFof_arguments(c *Fof_argumentsContext)

	// EnterFof_term is called when entering the fof_term production.
	EnterFof_term(c *Fof_termContext)

	// EnterFof_function_term is called when entering the fof_function_term production.
	EnterFof_function_term(c *Fof_function_termContext)

	// EnterTff_conditional_term is called when entering the tff_conditional_term production.
	EnterTff_conditional_term(c *Tff_conditional_termContext)

	// EnterTff_let_term is called when entering the tff_let_term production.
	EnterTff_let_term(c *Tff_let_termContext)

	// EnterTff_tuple_term is called when entering the tff_tuple_term production.
	EnterTff_tuple_term(c *Tff_tuple_termContext)

	// EnterFof_sequent is called when entering the fof_sequent production.
	EnterFof_sequent(c *Fof_sequentContext)

	// EnterFof_formula_tuple is called when entering the fof_formula_tuple production.
	EnterFof_formula_tuple(c *Fof_formula_tupleContext)

	// EnterFof_formula_tuple_list is called when entering the fof_formula_tuple_list production.
	EnterFof_formula_tuple_list(c *Fof_formula_tuple_listContext)

	// EnterCnf_formula is called when entering the cnf_formula production.
	EnterCnf_formula(c *Cnf_formulaContext)

	// EnterCnf_disjunction is called when entering the cnf_disjunction production.
	EnterCnf_disjunction(c *Cnf_disjunctionContext)

	// EnterCnf_literal is called when entering the cnf_literal production.
	EnterCnf_literal(c *Cnf_literalContext)

	// EnterThf_quantifier is called when entering the thf_quantifier production.
	EnterThf_quantifier(c *Thf_quantifierContext)

	// EnterTh0_quantifier is called when entering the th0_quantifier production.
	EnterTh0_quantifier(c *Th0_quantifierContext)

	// EnterTh1_quantifier is called when entering the th1_quantifier production.
	EnterTh1_quantifier(c *Th1_quantifierContext)

	// EnterThf_pair_connective is called when entering the thf_pair_connective production.
	EnterThf_pair_connective(c *Thf_pair_connectiveContext)

	// EnterThf_unary_connective is called when entering the thf_unary_connective production.
	EnterThf_unary_connective(c *Thf_unary_connectiveContext)

	// EnterTh1_unary_connective is called when entering the th1_unary_connective production.
	EnterTh1_unary_connective(c *Th1_unary_connectiveContext)

	// EnterTff_pair_connective is called when entering the tff_pair_connective production.
	EnterTff_pair_connective(c *Tff_pair_connectiveContext)

	// EnterFof_quantifier is called when entering the fof_quantifier production.
	EnterFof_quantifier(c *Fof_quantifierContext)

	// EnterBinary_connective is called when entering the binary_connective production.
	EnterBinary_connective(c *Binary_connectiveContext)

	// EnterAssoc_connective is called when entering the assoc_connective production.
	EnterAssoc_connective(c *Assoc_connectiveContext)

	// EnterUnary_connective is called when entering the unary_connective production.
	EnterUnary_connective(c *Unary_connectiveContext)

	// EnterType_constant is called when entering the type_constant production.
	EnterType_constant(c *Type_constantContext)

	// EnterType_functor is called when entering the type_functor production.
	EnterType_functor(c *Type_functorContext)

	// EnterDefined_type is called when entering the defined_type production.
	EnterDefined_type(c *Defined_typeContext)

	// EnterSystem_type is called when entering the system_type production.
	EnterSystem_type(c *System_typeContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterUntyped_atom is called when entering the untyped_atom production.
	EnterUntyped_atom(c *Untyped_atomContext)

	// EnterDefined_proposition is called when entering the defined_proposition production.
	EnterDefined_proposition(c *Defined_propositionContext)

	// EnterDefined_predicate is called when entering the defined_predicate production.
	EnterDefined_predicate(c *Defined_predicateContext)

	// EnterDefined_infix_pred is called when entering the defined_infix_pred production.
	EnterDefined_infix_pred(c *Defined_infix_predContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterFunctor is called when entering the functor production.
	EnterFunctor(c *FunctorContext)

	// EnterSystem_constant is called when entering the system_constant production.
	EnterSystem_constant(c *System_constantContext)

	// EnterSystem_functor is called when entering the system_functor production.
	EnterSystem_functor(c *System_functorContext)

	// EnterDefined_constant is called when entering the defined_constant production.
	EnterDefined_constant(c *Defined_constantContext)

	// EnterDefined_functor is called when entering the defined_functor production.
	EnterDefined_functor(c *Defined_functorContext)

	// EnterDefined_term is called when entering the defined_term production.
	EnterDefined_term(c *Defined_termContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterSource is called when entering the source production.
	EnterSource(c *SourceContext)

	// EnterSources is called when entering the sources production.
	EnterSources(c *SourcesContext)

	// EnterDag_source is called when entering the dag_source production.
	EnterDag_source(c *Dag_sourceContext)

	// EnterInference_record is called when entering the inference_record production.
	EnterInference_record(c *Inference_recordContext)

	// EnterInference_rule is called when entering the inference_rule production.
	EnterInference_rule(c *Inference_ruleContext)

	// EnterInference_parents is called when entering the inference_parents production.
	EnterInference_parents(c *Inference_parentsContext)

	// EnterParent_list is called when entering the parent_list production.
	EnterParent_list(c *Parent_listContext)

	// EnterParent_info is called when entering the parent_info production.
	EnterParent_info(c *Parent_infoContext)

	// EnterParent_details is called when entering the parent_details production.
	EnterParent_details(c *Parent_detailsContext)

	// EnterInternal_source is called when entering the internal_source production.
	EnterInternal_source(c *Internal_sourceContext)

	// EnterIntro_type is called when entering the intro_type production.
	EnterIntro_type(c *Intro_typeContext)

	// EnterExternal_source is called when entering the external_source production.
	EnterExternal_source(c *External_sourceContext)

	// EnterFile_source is called when entering the file_source production.
	EnterFile_source(c *File_sourceContext)

	// EnterFile_info is called when entering the file_info production.
	EnterFile_info(c *File_infoContext)

	// EnterTheory is called when entering the theory production.
	EnterTheory(c *TheoryContext)

	// EnterTheory_name is called when entering the theory_name production.
	EnterTheory_name(c *Theory_nameContext)

	// EnterCreator_source is called when entering the creator_source production.
	EnterCreator_source(c *Creator_sourceContext)

	// EnterCreator_name is called when entering the creator_name production.
	EnterCreator_name(c *Creator_nameContext)

	// EnterOptional_info is called when entering the optional_info production.
	EnterOptional_info(c *Optional_infoContext)

	// EnterUseful_info is called when entering the useful_info production.
	EnterUseful_info(c *Useful_infoContext)

	// EnterInfo_items is called when entering the info_items production.
	EnterInfo_items(c *Info_itemsContext)

	// EnterInfo_item is called when entering the info_item production.
	EnterInfo_item(c *Info_itemContext)

	// EnterFormula_item is called when entering the formula_item production.
	EnterFormula_item(c *Formula_itemContext)

	// EnterDescription_item is called when entering the description_item production.
	EnterDescription_item(c *Description_itemContext)

	// EnterIquote_item is called when entering the iquote_item production.
	EnterIquote_item(c *Iquote_itemContext)

	// EnterInference_item is called when entering the inference_item production.
	EnterInference_item(c *Inference_itemContext)

	// EnterInference_status is called when entering the inference_status production.
	EnterInference_status(c *Inference_statusContext)

	// EnterStatus_value is called when entering the status_value production.
	EnterStatus_value(c *Status_valueContext)

	// EnterInference_info is called when entering the inference_info production.
	EnterInference_info(c *Inference_infoContext)

	// EnterAssumptions_record is called when entering the assumptions_record production.
	EnterAssumptions_record(c *Assumptions_recordContext)

	// EnterRefutation is called when entering the refutation production.
	EnterRefutation(c *RefutationContext)

	// EnterNew_symbol_record is called when entering the new_symbol_record production.
	EnterNew_symbol_record(c *New_symbol_recordContext)

	// EnterNew_symbol_list is called when entering the new_symbol_list production.
	EnterNew_symbol_list(c *New_symbol_listContext)

	// EnterPrincipal_symbol is called when entering the principal_symbol production.
	EnterPrincipal_symbol(c *Principal_symbolContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterFormula_selection is called when entering the formula_selection production.
	EnterFormula_selection(c *Formula_selectionContext)

	// EnterName_list is called when entering the name_list production.
	EnterName_list(c *Name_listContext)

	// EnterGeneral_term is called when entering the general_term production.
	EnterGeneral_term(c *General_termContext)

	// EnterGeneral_data is called when entering the general_data production.
	EnterGeneral_data(c *General_dataContext)

	// EnterGeneral_function is called when entering the general_function production.
	EnterGeneral_function(c *General_functionContext)

	// EnterFormula_data is called when entering the formula_data production.
	EnterFormula_data(c *Formula_dataContext)

	// EnterGeneral_list is called when entering the general_list production.
	EnterGeneral_list(c *General_listContext)

	// EnterGeneral_terms is called when entering the general_terms production.
	EnterGeneral_terms(c *General_termsContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterAtomic_word is called when entering the atomic_word production.
	EnterAtomic_word(c *Atomic_wordContext)

	// EnterAtomic_defined_word is called when entering the atomic_defined_word production.
	EnterAtomic_defined_word(c *Atomic_defined_wordContext)

	// EnterAtomic_system_word is called when entering the atomic_system_word production.
	EnterAtomic_system_word(c *Atomic_system_wordContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterFile_name is called when entering the file_name production.
	EnterFile_name(c *File_nameContext)

	// ExitTptp_file is called when exiting the tptp_file production.
	ExitTptp_file(c *Tptp_fileContext)

	// ExitTptp_input is called when exiting the tptp_input production.
	ExitTptp_input(c *Tptp_inputContext)

	// ExitAnnotated_formula is called when exiting the annotated_formula production.
	ExitAnnotated_formula(c *Annotated_formulaContext)

	// ExitTpi_annotated is called when exiting the tpi_annotated production.
	ExitTpi_annotated(c *Tpi_annotatedContext)

	// ExitTpi_formula is called when exiting the tpi_formula production.
	ExitTpi_formula(c *Tpi_formulaContext)

	// ExitThf_annotated is called when exiting the thf_annotated production.
	ExitThf_annotated(c *Thf_annotatedContext)

	// ExitTfx_annotated is called when exiting the tfx_annotated production.
	ExitTfx_annotated(c *Tfx_annotatedContext)

	// ExitTff_annotated is called when exiting the tff_annotated production.
	ExitTff_annotated(c *Tff_annotatedContext)

	// ExitTcf_annotated is called when exiting the tcf_annotated production.
	ExitTcf_annotated(c *Tcf_annotatedContext)

	// ExitFof_annotated is called when exiting the fof_annotated production.
	ExitFof_annotated(c *Fof_annotatedContext)

	// ExitCnf_annotated is called when exiting the cnf_annotated production.
	ExitCnf_annotated(c *Cnf_annotatedContext)

	// ExitAnnotations is called when exiting the annotations production.
	ExitAnnotations(c *AnnotationsContext)

	// ExitFormula_role is called when exiting the formula_role production.
	ExitFormula_role(c *Formula_roleContext)

	// ExitThf_formula is called when exiting the thf_formula production.
	ExitThf_formula(c *Thf_formulaContext)

	// ExitThf_logic_formula is called when exiting the thf_logic_formula production.
	ExitThf_logic_formula(c *Thf_logic_formulaContext)

	// ExitThf_binary_formula is called when exiting the thf_binary_formula production.
	ExitThf_binary_formula(c *Thf_binary_formulaContext)

	// ExitThf_binary_pair is called when exiting the thf_binary_pair production.
	ExitThf_binary_pair(c *Thf_binary_pairContext)

	// ExitThf_binary_tuple is called when exiting the thf_binary_tuple production.
	ExitThf_binary_tuple(c *Thf_binary_tupleContext)

	// ExitThf_or_formula is called when exiting the thf_or_formula production.
	ExitThf_or_formula(c *Thf_or_formulaContext)

	// ExitThf_and_formula is called when exiting the thf_and_formula production.
	ExitThf_and_formula(c *Thf_and_formulaContext)

	// ExitThf_apply_formula is called when exiting the thf_apply_formula production.
	ExitThf_apply_formula(c *Thf_apply_formulaContext)

	// ExitThf_unitary_formula is called when exiting the thf_unitary_formula production.
	ExitThf_unitary_formula(c *Thf_unitary_formulaContext)

	// ExitThf_quantified_formula is called when exiting the thf_quantified_formula production.
	ExitThf_quantified_formula(c *Thf_quantified_formulaContext)

	// ExitThf_quantification is called when exiting the thf_quantification production.
	ExitThf_quantification(c *Thf_quantificationContext)

	// ExitThf_variable_list is called when exiting the thf_variable_list production.
	ExitThf_variable_list(c *Thf_variable_listContext)

	// ExitThf_variable is called when exiting the thf_variable production.
	ExitThf_variable(c *Thf_variableContext)

	// ExitThf_typed_variable is called when exiting the thf_typed_variable production.
	ExitThf_typed_variable(c *Thf_typed_variableContext)

	// ExitThf_unary_formula is called when exiting the thf_unary_formula production.
	ExitThf_unary_formula(c *Thf_unary_formulaContext)

	// ExitThf_atom is called when exiting the thf_atom production.
	ExitThf_atom(c *Thf_atomContext)

	// ExitThf_function is called when exiting the thf_function production.
	ExitThf_function(c *Thf_functionContext)

	// ExitThf_conn_term is called when exiting the thf_conn_term production.
	ExitThf_conn_term(c *Thf_conn_termContext)

	// ExitThf_conditional is called when exiting the thf_conditional production.
	ExitThf_conditional(c *Thf_conditionalContext)

	// ExitThf_let is called when exiting the thf_let production.
	ExitThf_let(c *Thf_letContext)

	// ExitThf_arguments is called when exiting the thf_arguments production.
	ExitThf_arguments(c *Thf_argumentsContext)

	// ExitThf_type_formula is called when exiting the thf_type_formula production.
	ExitThf_type_formula(c *Thf_type_formulaContext)

	// ExitThf_typeable_formula is called when exiting the thf_typeable_formula production.
	ExitThf_typeable_formula(c *Thf_typeable_formulaContext)

	// ExitThf_subtype is called when exiting the thf_subtype production.
	ExitThf_subtype(c *Thf_subtypeContext)

	// ExitThf_top_level_type is called when exiting the thf_top_level_type production.
	ExitThf_top_level_type(c *Thf_top_level_typeContext)

	// ExitThf_unitary_type is called when exiting the thf_unitary_type production.
	ExitThf_unitary_type(c *Thf_unitary_typeContext)

	// ExitThf_apply_type is called when exiting the thf_apply_type production.
	ExitThf_apply_type(c *Thf_apply_typeContext)

	// ExitThf_binary_type is called when exiting the thf_binary_type production.
	ExitThf_binary_type(c *Thf_binary_typeContext)

	// ExitThf_mapping_type is called when exiting the thf_mapping_type production.
	ExitThf_mapping_type(c *Thf_mapping_typeContext)

	// ExitThf_xprod_type is called when exiting the thf_xprod_type production.
	ExitThf_xprod_type(c *Thf_xprod_typeContext)

	// ExitThf_union_type is called when exiting the thf_union_type production.
	ExitThf_union_type(c *Thf_union_typeContext)

	// ExitThf_sequent is called when exiting the thf_sequent production.
	ExitThf_sequent(c *Thf_sequentContext)

	// ExitThf_tuple is called when exiting the thf_tuple production.
	ExitThf_tuple(c *Thf_tupleContext)

	// ExitThf_formula_list is called when exiting the thf_formula_list production.
	ExitThf_formula_list(c *Thf_formula_listContext)

	// ExitTfx_formula is called when exiting the tfx_formula production.
	ExitTfx_formula(c *Tfx_formulaContext)

	// ExitTfx_logic_formula is called when exiting the tfx_logic_formula production.
	ExitTfx_logic_formula(c *Tfx_logic_formulaContext)

	// ExitTff_formula is called when exiting the tff_formula production.
	ExitTff_formula(c *Tff_formulaContext)

	// ExitTff_logic_formula is called when exiting the tff_logic_formula production.
	ExitTff_logic_formula(c *Tff_logic_formulaContext)

	// ExitTff_binary_formula is called when exiting the tff_binary_formula production.
	ExitTff_binary_formula(c *Tff_binary_formulaContext)

	// ExitTff_binary_nonassoc is called when exiting the tff_binary_nonassoc production.
	ExitTff_binary_nonassoc(c *Tff_binary_nonassocContext)

	// ExitTff_binary_assoc is called when exiting the tff_binary_assoc production.
	ExitTff_binary_assoc(c *Tff_binary_assocContext)

	// ExitTff_or_formula is called when exiting the tff_or_formula production.
	ExitTff_or_formula(c *Tff_or_formulaContext)

	// ExitTff_and_formula is called when exiting the tff_and_formula production.
	ExitTff_and_formula(c *Tff_and_formulaContext)

	// ExitTff_unitary_formula is called when exiting the tff_unitary_formula production.
	ExitTff_unitary_formula(c *Tff_unitary_formulaContext)

	// ExitTff_quantified_formula is called when exiting the tff_quantified_formula production.
	ExitTff_quantified_formula(c *Tff_quantified_formulaContext)

	// ExitTff_variable_list is called when exiting the tff_variable_list production.
	ExitTff_variable_list(c *Tff_variable_listContext)

	// ExitTff_variable is called when exiting the tff_variable production.
	ExitTff_variable(c *Tff_variableContext)

	// ExitTff_typed_variable is called when exiting the tff_typed_variable production.
	ExitTff_typed_variable(c *Tff_typed_variableContext)

	// ExitTff_unary_formula is called when exiting the tff_unary_formula production.
	ExitTff_unary_formula(c *Tff_unary_formulaContext)

	// ExitTff_atomic_formula is called when exiting the tff_atomic_formula production.
	ExitTff_atomic_formula(c *Tff_atomic_formulaContext)

	// ExitTff_conditional is called when exiting the tff_conditional production.
	ExitTff_conditional(c *Tff_conditionalContext)

	// ExitTff_let is called when exiting the tff_let production.
	ExitTff_let(c *Tff_letContext)

	// ExitTff_let_term_defns is called when exiting the tff_let_term_defns production.
	ExitTff_let_term_defns(c *Tff_let_term_defnsContext)

	// ExitTff_let_term_list is called when exiting the tff_let_term_list production.
	ExitTff_let_term_list(c *Tff_let_term_listContext)

	// ExitTff_let_term_defn is called when exiting the tff_let_term_defn production.
	ExitTff_let_term_defn(c *Tff_let_term_defnContext)

	// ExitTff_let_term_binding is called when exiting the tff_let_term_binding production.
	ExitTff_let_term_binding(c *Tff_let_term_bindingContext)

	// ExitTff_let_formula_defns is called when exiting the tff_let_formula_defns production.
	ExitTff_let_formula_defns(c *Tff_let_formula_defnsContext)

	// ExitTff_let_formula_list is called when exiting the tff_let_formula_list production.
	ExitTff_let_formula_list(c *Tff_let_formula_listContext)

	// ExitTff_let_formula_defn is called when exiting the tff_let_formula_defn production.
	ExitTff_let_formula_defn(c *Tff_let_formula_defnContext)

	// ExitTff_let_formula_binding is called when exiting the tff_let_formula_binding production.
	ExitTff_let_formula_binding(c *Tff_let_formula_bindingContext)

	// ExitTff_sequent is called when exiting the tff_sequent production.
	ExitTff_sequent(c *Tff_sequentContext)

	// ExitTff_formula_tuple is called when exiting the tff_formula_tuple production.
	ExitTff_formula_tuple(c *Tff_formula_tupleContext)

	// ExitTff_formula_tuple_list is called when exiting the tff_formula_tuple_list production.
	ExitTff_formula_tuple_list(c *Tff_formula_tuple_listContext)

	// ExitTff_typed_atom is called when exiting the tff_typed_atom production.
	ExitTff_typed_atom(c *Tff_typed_atomContext)

	// ExitTff_subtype is called when exiting the tff_subtype production.
	ExitTff_subtype(c *Tff_subtypeContext)

	// ExitTff_top_level_type is called when exiting the tff_top_level_type production.
	ExitTff_top_level_type(c *Tff_top_level_typeContext)

	// ExitTf1_quantified_type is called when exiting the tf1_quantified_type production.
	ExitTf1_quantified_type(c *Tf1_quantified_typeContext)

	// ExitTff_monotype is called when exiting the tff_monotype production.
	ExitTff_monotype(c *Tff_monotypeContext)

	// ExitTff_unitary_type is called when exiting the tff_unitary_type production.
	ExitTff_unitary_type(c *Tff_unitary_typeContext)

	// ExitTff_atomic_type is called when exiting the tff_atomic_type production.
	ExitTff_atomic_type(c *Tff_atomic_typeContext)

	// ExitTff_type_arguments is called when exiting the tff_type_arguments production.
	ExitTff_type_arguments(c *Tff_type_argumentsContext)

	// ExitTff_mapping_type is called when exiting the tff_mapping_type production.
	ExitTff_mapping_type(c *Tff_mapping_typeContext)

	// ExitTff_xprod_type is called when exiting the tff_xprod_type production.
	ExitTff_xprod_type(c *Tff_xprod_typeContext)

	// ExitTcf_formula is called when exiting the tcf_formula production.
	ExitTcf_formula(c *Tcf_formulaContext)

	// ExitTcf_logic_formula is called when exiting the tcf_logic_formula production.
	ExitTcf_logic_formula(c *Tcf_logic_formulaContext)

	// ExitTcf_quantified_formula is called when exiting the tcf_quantified_formula production.
	ExitTcf_quantified_formula(c *Tcf_quantified_formulaContext)

	// ExitFof_formula is called when exiting the fof_formula production.
	ExitFof_formula(c *Fof_formulaContext)

	// ExitFof_logic_formula is called when exiting the fof_logic_formula production.
	ExitFof_logic_formula(c *Fof_logic_formulaContext)

	// ExitFof_binary_formula is called when exiting the fof_binary_formula production.
	ExitFof_binary_formula(c *Fof_binary_formulaContext)

	// ExitFof_binary_nonassoc is called when exiting the fof_binary_nonassoc production.
	ExitFof_binary_nonassoc(c *Fof_binary_nonassocContext)

	// ExitFof_binary_assoc is called when exiting the fof_binary_assoc production.
	ExitFof_binary_assoc(c *Fof_binary_assocContext)

	// ExitFof_or_formula is called when exiting the fof_or_formula production.
	ExitFof_or_formula(c *Fof_or_formulaContext)

	// ExitFof_and_formula is called when exiting the fof_and_formula production.
	ExitFof_and_formula(c *Fof_and_formulaContext)

	// ExitFof_unitary_formula is called when exiting the fof_unitary_formula production.
	ExitFof_unitary_formula(c *Fof_unitary_formulaContext)

	// ExitFof_quantified_formula is called when exiting the fof_quantified_formula production.
	ExitFof_quantified_formula(c *Fof_quantified_formulaContext)

	// ExitFof_variable_list is called when exiting the fof_variable_list production.
	ExitFof_variable_list(c *Fof_variable_listContext)

	// ExitFof_unary_formula is called when exiting the fof_unary_formula production.
	ExitFof_unary_formula(c *Fof_unary_formulaContext)

	// ExitFof_infix_unary is called when exiting the fof_infix_unary production.
	ExitFof_infix_unary(c *Fof_infix_unaryContext)

	// ExitFof_atomic_formula is called when exiting the fof_atomic_formula production.
	ExitFof_atomic_formula(c *Fof_atomic_formulaContext)

	// ExitFof_plain_atomic_formula is called when exiting the fof_plain_atomic_formula production.
	ExitFof_plain_atomic_formula(c *Fof_plain_atomic_formulaContext)

	// ExitFof_defined_atomic_formula is called when exiting the fof_defined_atomic_formula production.
	ExitFof_defined_atomic_formula(c *Fof_defined_atomic_formulaContext)

	// ExitFof_defined_plain_formula is called when exiting the fof_defined_plain_formula production.
	ExitFof_defined_plain_formula(c *Fof_defined_plain_formulaContext)

	// ExitFof_defined_infix_formula is called when exiting the fof_defined_infix_formula production.
	ExitFof_defined_infix_formula(c *Fof_defined_infix_formulaContext)

	// ExitFof_system_atomic_formula is called when exiting the fof_system_atomic_formula production.
	ExitFof_system_atomic_formula(c *Fof_system_atomic_formulaContext)

	// ExitFof_plain_term is called when exiting the fof_plain_term production.
	ExitFof_plain_term(c *Fof_plain_termContext)

	// ExitFof_defined_term is called when exiting the fof_defined_term production.
	ExitFof_defined_term(c *Fof_defined_termContext)

	// ExitFof_defined_atomic_term is called when exiting the fof_defined_atomic_term production.
	ExitFof_defined_atomic_term(c *Fof_defined_atomic_termContext)

	// ExitFof_defined_plain_term is called when exiting the fof_defined_plain_term production.
	ExitFof_defined_plain_term(c *Fof_defined_plain_termContext)

	// ExitFof_system_term is called when exiting the fof_system_term production.
	ExitFof_system_term(c *Fof_system_termContext)

	// ExitFof_arguments is called when exiting the fof_arguments production.
	ExitFof_arguments(c *Fof_argumentsContext)

	// ExitFof_term is called when exiting the fof_term production.
	ExitFof_term(c *Fof_termContext)

	// ExitFof_function_term is called when exiting the fof_function_term production.
	ExitFof_function_term(c *Fof_function_termContext)

	// ExitTff_conditional_term is called when exiting the tff_conditional_term production.
	ExitTff_conditional_term(c *Tff_conditional_termContext)

	// ExitTff_let_term is called when exiting the tff_let_term production.
	ExitTff_let_term(c *Tff_let_termContext)

	// ExitTff_tuple_term is called when exiting the tff_tuple_term production.
	ExitTff_tuple_term(c *Tff_tuple_termContext)

	// ExitFof_sequent is called when exiting the fof_sequent production.
	ExitFof_sequent(c *Fof_sequentContext)

	// ExitFof_formula_tuple is called when exiting the fof_formula_tuple production.
	ExitFof_formula_tuple(c *Fof_formula_tupleContext)

	// ExitFof_formula_tuple_list is called when exiting the fof_formula_tuple_list production.
	ExitFof_formula_tuple_list(c *Fof_formula_tuple_listContext)

	// ExitCnf_formula is called when exiting the cnf_formula production.
	ExitCnf_formula(c *Cnf_formulaContext)

	// ExitCnf_disjunction is called when exiting the cnf_disjunction production.
	ExitCnf_disjunction(c *Cnf_disjunctionContext)

	// ExitCnf_literal is called when exiting the cnf_literal production.
	ExitCnf_literal(c *Cnf_literalContext)

	// ExitThf_quantifier is called when exiting the thf_quantifier production.
	ExitThf_quantifier(c *Thf_quantifierContext)

	// ExitTh0_quantifier is called when exiting the th0_quantifier production.
	ExitTh0_quantifier(c *Th0_quantifierContext)

	// ExitTh1_quantifier is called when exiting the th1_quantifier production.
	ExitTh1_quantifier(c *Th1_quantifierContext)

	// ExitThf_pair_connective is called when exiting the thf_pair_connective production.
	ExitThf_pair_connective(c *Thf_pair_connectiveContext)

	// ExitThf_unary_connective is called when exiting the thf_unary_connective production.
	ExitThf_unary_connective(c *Thf_unary_connectiveContext)

	// ExitTh1_unary_connective is called when exiting the th1_unary_connective production.
	ExitTh1_unary_connective(c *Th1_unary_connectiveContext)

	// ExitTff_pair_connective is called when exiting the tff_pair_connective production.
	ExitTff_pair_connective(c *Tff_pair_connectiveContext)

	// ExitFof_quantifier is called when exiting the fof_quantifier production.
	ExitFof_quantifier(c *Fof_quantifierContext)

	// ExitBinary_connective is called when exiting the binary_connective production.
	ExitBinary_connective(c *Binary_connectiveContext)

	// ExitAssoc_connective is called when exiting the assoc_connective production.
	ExitAssoc_connective(c *Assoc_connectiveContext)

	// ExitUnary_connective is called when exiting the unary_connective production.
	ExitUnary_connective(c *Unary_connectiveContext)

	// ExitType_constant is called when exiting the type_constant production.
	ExitType_constant(c *Type_constantContext)

	// ExitType_functor is called when exiting the type_functor production.
	ExitType_functor(c *Type_functorContext)

	// ExitDefined_type is called when exiting the defined_type production.
	ExitDefined_type(c *Defined_typeContext)

	// ExitSystem_type is called when exiting the system_type production.
	ExitSystem_type(c *System_typeContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitUntyped_atom is called when exiting the untyped_atom production.
	ExitUntyped_atom(c *Untyped_atomContext)

	// ExitDefined_proposition is called when exiting the defined_proposition production.
	ExitDefined_proposition(c *Defined_propositionContext)

	// ExitDefined_predicate is called when exiting the defined_predicate production.
	ExitDefined_predicate(c *Defined_predicateContext)

	// ExitDefined_infix_pred is called when exiting the defined_infix_pred production.
	ExitDefined_infix_pred(c *Defined_infix_predContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitFunctor is called when exiting the functor production.
	ExitFunctor(c *FunctorContext)

	// ExitSystem_constant is called when exiting the system_constant production.
	ExitSystem_constant(c *System_constantContext)

	// ExitSystem_functor is called when exiting the system_functor production.
	ExitSystem_functor(c *System_functorContext)

	// ExitDefined_constant is called when exiting the defined_constant production.
	ExitDefined_constant(c *Defined_constantContext)

	// ExitDefined_functor is called when exiting the defined_functor production.
	ExitDefined_functor(c *Defined_functorContext)

	// ExitDefined_term is called when exiting the defined_term production.
	ExitDefined_term(c *Defined_termContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)

	// ExitSources is called when exiting the sources production.
	ExitSources(c *SourcesContext)

	// ExitDag_source is called when exiting the dag_source production.
	ExitDag_source(c *Dag_sourceContext)

	// ExitInference_record is called when exiting the inference_record production.
	ExitInference_record(c *Inference_recordContext)

	// ExitInference_rule is called when exiting the inference_rule production.
	ExitInference_rule(c *Inference_ruleContext)

	// ExitInference_parents is called when exiting the inference_parents production.
	ExitInference_parents(c *Inference_parentsContext)

	// ExitParent_list is called when exiting the parent_list production.
	ExitParent_list(c *Parent_listContext)

	// ExitParent_info is called when exiting the parent_info production.
	ExitParent_info(c *Parent_infoContext)

	// ExitParent_details is called when exiting the parent_details production.
	ExitParent_details(c *Parent_detailsContext)

	// ExitInternal_source is called when exiting the internal_source production.
	ExitInternal_source(c *Internal_sourceContext)

	// ExitIntro_type is called when exiting the intro_type production.
	ExitIntro_type(c *Intro_typeContext)

	// ExitExternal_source is called when exiting the external_source production.
	ExitExternal_source(c *External_sourceContext)

	// ExitFile_source is called when exiting the file_source production.
	ExitFile_source(c *File_sourceContext)

	// ExitFile_info is called when exiting the file_info production.
	ExitFile_info(c *File_infoContext)

	// ExitTheory is called when exiting the theory production.
	ExitTheory(c *TheoryContext)

	// ExitTheory_name is called when exiting the theory_name production.
	ExitTheory_name(c *Theory_nameContext)

	// ExitCreator_source is called when exiting the creator_source production.
	ExitCreator_source(c *Creator_sourceContext)

	// ExitCreator_name is called when exiting the creator_name production.
	ExitCreator_name(c *Creator_nameContext)

	// ExitOptional_info is called when exiting the optional_info production.
	ExitOptional_info(c *Optional_infoContext)

	// ExitUseful_info is called when exiting the useful_info production.
	ExitUseful_info(c *Useful_infoContext)

	// ExitInfo_items is called when exiting the info_items production.
	ExitInfo_items(c *Info_itemsContext)

	// ExitInfo_item is called when exiting the info_item production.
	ExitInfo_item(c *Info_itemContext)

	// ExitFormula_item is called when exiting the formula_item production.
	ExitFormula_item(c *Formula_itemContext)

	// ExitDescription_item is called when exiting the description_item production.
	ExitDescription_item(c *Description_itemContext)

	// ExitIquote_item is called when exiting the iquote_item production.
	ExitIquote_item(c *Iquote_itemContext)

	// ExitInference_item is called when exiting the inference_item production.
	ExitInference_item(c *Inference_itemContext)

	// ExitInference_status is called when exiting the inference_status production.
	ExitInference_status(c *Inference_statusContext)

	// ExitStatus_value is called when exiting the status_value production.
	ExitStatus_value(c *Status_valueContext)

	// ExitInference_info is called when exiting the inference_info production.
	ExitInference_info(c *Inference_infoContext)

	// ExitAssumptions_record is called when exiting the assumptions_record production.
	ExitAssumptions_record(c *Assumptions_recordContext)

	// ExitRefutation is called when exiting the refutation production.
	ExitRefutation(c *RefutationContext)

	// ExitNew_symbol_record is called when exiting the new_symbol_record production.
	ExitNew_symbol_record(c *New_symbol_recordContext)

	// ExitNew_symbol_list is called when exiting the new_symbol_list production.
	ExitNew_symbol_list(c *New_symbol_listContext)

	// ExitPrincipal_symbol is called when exiting the principal_symbol production.
	ExitPrincipal_symbol(c *Principal_symbolContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitFormula_selection is called when exiting the formula_selection production.
	ExitFormula_selection(c *Formula_selectionContext)

	// ExitName_list is called when exiting the name_list production.
	ExitName_list(c *Name_listContext)

	// ExitGeneral_term is called when exiting the general_term production.
	ExitGeneral_term(c *General_termContext)

	// ExitGeneral_data is called when exiting the general_data production.
	ExitGeneral_data(c *General_dataContext)

	// ExitGeneral_function is called when exiting the general_function production.
	ExitGeneral_function(c *General_functionContext)

	// ExitFormula_data is called when exiting the formula_data production.
	ExitFormula_data(c *Formula_dataContext)

	// ExitGeneral_list is called when exiting the general_list production.
	ExitGeneral_list(c *General_listContext)

	// ExitGeneral_terms is called when exiting the general_terms production.
	ExitGeneral_terms(c *General_termsContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitAtomic_word is called when exiting the atomic_word production.
	ExitAtomic_word(c *Atomic_wordContext)

	// ExitAtomic_defined_word is called when exiting the atomic_defined_word production.
	ExitAtomic_defined_word(c *Atomic_defined_wordContext)

	// ExitAtomic_system_word is called when exiting the atomic_system_word production.
	ExitAtomic_system_word(c *Atomic_system_wordContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitFile_name is called when exiting the file_name production.
	ExitFile_name(c *File_nameContext)
}
