% Test file for best-first search program

:- use_module('fig12_3.pl').
:- use_module('fig12_6.pl').
:- use_module(library(plunit)).

% Set dynamic heuristic predicate with cuts
set_heuristic(1) :- !,
    retractall(current_heuristic(_)),
    assert(current_heuristic(heuristic1)).
set_heuristic(2) :- !,
    retractall(current_heuristic(_)),
    assert(current_heuristic(heuristic2)).
set_heuristic(3) :- !,
    retractall(current_heuristic(_)),
    assert(current_heuristic(heuristic3)).

:- begin_tests(hw6).

test(heuristic1, [setup(set_heuristic(1))]) :-
    start1(Pos),
    bestfirst(Pos, Sol),
    nodes_count(Count),
    format('Heuristic1 generated ~w nodes~n', [Count]),
    assertion(Count > 0),
    !.

test(heuristic2, [setup(set_heuristic(2))]) :-
    start1(Pos),
    bestfirst(Pos, Sol),
    nodes_count(Count),
    format('Heuristic2 generated ~w nodes~n', [Count]),
    assertion(Count > 0),
    !.

test(heuristic3, [setup(set_heuristic(3))]) :-
    start1(Pos),
    bestfirst(Pos, Sol),
    nodes_count(Count),
    format('Heuristic3 generated ~w nodes~n', [Count]),
    assertion(Count > 0),
    !.

:- end_tests(hw6).
