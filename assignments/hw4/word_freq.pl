:- use_module(library(readutil)).
% Write source code here

word_frequencies(File, StopWordsFile):-
    % 1. 讀文字
    read_file_to_string(File, Raw, []),
    % 2. 過濾小寫
    filter_chars_and_normalize(Raw, Clean),
    % 3. 切字
    scan(Clean, Words0),
    % 4. 讀停用詞
    read_stop_words(StopWordsFile, StopWords),
    % 5. 濾停用詞
    remove_stop_words(Words0, StopWords, Words),
    % 6. 統計次數
    frequencies(Words, Freq0),
    % 7. 排序
    sorted(Freq0, FreqSorted),
    % 8. 取前 25
    take(25, FreqSorted, Top25),
    % 9. 列印
    print_freqs(Top25).

% 取前 N 個
take(0, _, []) :- !.
take(_, [], []).
take(N, [H|T], [H|R]) :-
    N > 0,
    N1 is N - 1,
    take(N1, T, R).

% 列印清單
print_freqs([]).
print_freqs([W-N|T]) :-
    format('~w - ~d~n', [W, N]),
    print_freqs(T).


read_stop_words(File, StopWords) :-
    read_file_to_string(File, Text, []),
    split_string(Text, ",", "\n\t ", Words0),
    maplist(string_lower, Words0, Words1),
    Alphabet = ["a","b","c","d","e","f","g","h","i",
                "j","k","l","m","n","o","p","q","r",
                "s","t","u","v","w","x","y","z"],
    append(Words1, Alphabet, StopWords).

:- use_module(library(pcre)).

filter_chars_and_normalize(Text, Filtered) :-
    string_lower(Text, Lower),
    re_replace('[^a-z]+'/g, ' ', Lower, Filtered ).


scan(Text, WordList) :-
    split_string(Text, " ", "\n\t", Words),
    exclude(=( ""), Words, WordList).

remove_stop_words(WordList, StopWords, FilteredWordList):-
    subtract(WordList, StopWords, FilteredWordList).



frequencies(WordList, WordFreq) :-
    frequencies(WordList, [], WordFreq).

frequencies([], Acc, Acc).
frequencies([Word|Rest], Acc, WordFreq) :-
    % Convert Word from string to atom
    atom_string(WordAtom, Word),
    (   select(WordAtom-Count, Acc, RestAcc) ->
        NewCount is Count + 1,
        Acc1 = [WordAtom-NewCount|RestAcc]
    ;   
        Acc1 = [WordAtom-1|Acc]
    ),
    frequencies(Rest, Acc1, WordFreq).

sorted(WordList, SortedWordList) :-
    predsort(compare_freq, WordList, SortedWordList).

compare_freq(Delta, _-N1, _-N2) :-
    compare(Delta, N2, N1).
