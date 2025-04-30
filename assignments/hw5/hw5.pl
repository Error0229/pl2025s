% Class hierarchy definitions
class(lane).
class(stage).
class(swimlane).

% Define abstract and concrete classes
abstract_class(lane).
concrete_class(stage).
concrete_class(swimlane).

% Subclass relationships
subclass_of(stage, lane).
subclass_of(swimlane, lane).

% Dynamic predicates for instance management
:- dynamic instance_of/2.
:- dynamic child_of/2.

% Helper to check if a class is valid (either concrete or abstract)
valid_class(Class) :- 
    class(Class).

% Helper to check if an instance exists
instance_exists(Instance) :-
    instance_of(Instance, _).

% Create a new instance of a class
new_instance(Instance, Class) :-
    valid_class(Class),
    concrete_class(Class),
    assertz(instance_of(Instance, Class)), true.

% Get the class of an instance
get_class(Instance, Class) :-
    instance_of(Instance, Class).

% Check if two instances are of the same class
same_class(Instance1, Instance2) :-
    get_class(Instance1, ClassA),
    get_class(Instance2, ClassA).

% Add a child to a parent instance
add_child(Child, Parent) :-
    instance_exists(Child),
    instance_exists(Parent),
    \+ is_descendant(Parent, Child),  % Parent is not a descendant of Child
    \+ is_descendant(Child, Parent),  % Child is not a descendant of Parent
    \+ has_other_parent(Child),       % Child doesn't have another parent
    (has_children(Parent) ->          % If parent has children
        children(Parent, ExistingChildren),
        member(FirstChild, ExistingChildren),
        same_class(FirstChild, Child) % New child must be same class as existing children
        ;
        true                         % If no children, any valid instance can be added
    ),
    assertz(child_of(Child, Parent)), !.

add_child(Child, Parent) :-
    instance_exists(Child),
    instance_exists(Parent),
    (is_descendant(Parent, Child) ->
        writeln('Child not added: cannot add a descendant as a new child')
    ;
    is_descendant(Child, Parent) ->
        writeln('Child not added: cannot add an ancestor as a new child')
    ;
    has_other_parent(Child) ->
        writeln('Child not added: cannot add a child that is already a child of another parent')
    ;
    has_children(Parent),
    children(Parent, ExistingChildren),
    member(FirstChild, ExistingChildren),
    \+ same_class(FirstChild, Child) ->
        writeln('Child not added: cannot add a child with different class')
    ;
        true
    ), !.

% Check if an instance has any children
has_children(Instance) :-
    child_of(_, Instance).

% Get immediate children of an instance
children(Instance, Children) :-
    findall(Child, child_of(Child, Instance), Children), !.

% Check if Child is a descendant of Ancestor
is_descendant(Ancestor, Descendant) :-
    child_of(Descendant, Ancestor).
is_descendant(Ancestor, Descendant) :-
    child_of(Intermediate, Ancestor),
    is_descendant(Intermediate, Descendant).

% Get all descendants of an instance
descendants(Instance, Descendants) :-
    findall(D, is_descendant(Instance, D), Descendants), !.

% Check if an instance has a parent
has_other_parent(Child) :-
    child_of(Child, _).

% Validate that all children of an instance are of the same concrete class
valid_children(Instance) :-
    \+ has_children(Instance), !.
valid_children(Instance) :-
    children(Instance, Children),
    Children = [FirstChild|Rest],
    get_class(FirstChild, FirstClass),
    forall(member(OtherChild, Rest), 
           (get_class(OtherChild, OtherClass), FirstClass = OtherClass)), !.
