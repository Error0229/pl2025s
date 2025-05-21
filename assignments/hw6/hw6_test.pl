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
    write('Testing heuristic1:'), nl,
    flush_output,
    start1(Pos), 
    bestfirst(Pos, _),
    nodes_count(Count1),
    write('Start1 generated '), write(Count1), write(' nodes'), nl,
    start2(Pos2),
    bestfirst(Pos2, _),
    nodes_count(Count2),
    write('Start2 generated '), write(Count2), write(' nodes'), nl,
    start3(Pos3),
    bestfirst(Pos3, _),
    nodes_count(Count3),
    write('Start3 generated '), write(Count3), write(' nodes'), nl,
    flush_output,
    !.

test(heuristic2, [setup(set_heuristic(2))]) :-
    write('Testing heuristic2:'), nl,
    flush_output,
    start1(Pos),
    bestfirst(Pos, _),
    nodes_count(Count1),
    write('Start1 generated '), write(Count1), write(' nodes'), nl,
    start2(Pos2),
    bestfirst(Pos2, _),
    nodes_count(Count2),
    write('Start2 generated '), write(Count2), write(' nodes'), nl,
    start3(Pos3),
    bestfirst(Pos3, _),
    nodes_count(Count3),
    write('Start3 generated '), write(Count3), write(' nodes'), nl,
    flush_output,
    !.

test(heuristic3, [setup(set_heuristic(3))]) :-
    write('Testing heuristic3:'), nl,
    flush_output,
    start1(Pos),
    bestfirst(Pos, _),
    nodes_count(Count1),
    write('Start1 generated '), write(Count1), write(' nodes'), nl,
    start2(Pos2),
    bestfirst(Pos2, _),
    nodes_count(Count2),
    write('Start2 generated '), write(Count2), write(' nodes'), nl,
    start3(Pos3),
    bestfirst(Pos3, _),
    nodes_count(Count3),
    write('Start3 generated '), write(Count3), write(' nodes'), nl,
    flush_output,
    !.

:- end_tests(hw6).
