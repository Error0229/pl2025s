:- begin_tests(hw5).
:- use_module(library(plunit)).
:- consult(hw5).

% Clear all dynamic predicates before each test
clear_all :-
    retractall(instance_of(_, _)),
    retractall(child_of(_, _)).

test(class_hierarchy) :-
    class(lane),
    class(stage),
    class(swimlane),
    abstract_class(lane),
    concrete_class(stage),
    concrete_class(swimlane),
    subclass_of(stage, lane),
    subclass_of(swimlane, lane).

test(new_instance_concrete, [setup(clear_all)]) :-
    new_instance(i1, stage),
    new_instance(i2, swimlane),
    instance_of(i1, stage),
    instance_of(i2, swimlane).

test(new_instance_abstract, [fail, setup(clear_all)]) :-
    new_instance(i, lane).

test(add_child_basic, [setup(clear_all)]) :-
    new_instance(s1, stage),
    new_instance(l3, swimlane),
    add_child(l3, s1),
    children(s1, [l3]).

test(add_child_same_class, [setup(clear_all)]) :-
    new_instance(s1, stage),
    new_instance(l3, swimlane),
    new_instance(l4, swimlane),
    add_child(l3, s1),
    add_child(l4, s1),
    children(s1, Children),
    sort(Children, [l3, l4]).

test(add_child_different_class, [setup(clear_all)]) :-
    new_instance(s1, stage),
    new_instance(s2, stage),
    new_instance(l3, swimlane),
    add_child(l3, s1),
    add_child(s2, s1),
    children(s1, [l3]).

test(descendants, [setup(clear_all)]) :-
    new_instance(s1, stage),
    new_instance(l3, swimlane),
    new_instance(s2, stage),
    add_child(l3, s1),
    add_child(s2, l3),
    descendants(s1, Desc),
    sort(Desc, [l3, s2]).

test(no_ancestor_as_child, [setup(clear_all)]) :-
    new_instance(s1, stage),
    new_instance(l3, swimlane),
    new_instance(s2, stage),
    add_child(l3, s1),
    add_child(s2, l3),
    add_child(s1, s2),
    children(s2, []).

test(no_descendant_as_child, [setup(clear_all)]) :-
    new_instance(s1, stage),
    new_instance(l3, swimlane),
    new_instance(s2, stage),
    add_child(l3, s1),
    add_child(s2, l3),
    add_child(s2, s1),
    children(s1, [l3]).

test(no_multiple_parents, [setup(clear_all)]) :-
    new_instance(s1, stage),
    new_instance(s2, stage),
    new_instance(l5, swimlane),
    add_child(l5, s1),
    add_child(l5, s2),
    children(s2, []).

test(valid_children_check, [setup(clear_all)]) :-
    new_instance(s1, stage),
    new_instance(l3, swimlane),
    new_instance(l4, swimlane),
    add_child(l3, s1),
    add_child(l4, s1),
    valid_children(s1).

:- end_tests(hw5).
