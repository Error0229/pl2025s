:- use_module(library(plunit)).
:- use_module(word_freq).

:- begin_tests(word_freq).

% 1. filter_chars_and_normalize/2
test(filter_chars_and_normalize) :-
    filter_chars_and_normalize("Hello, WORLD!", Out),
    assertion(Out == "hello world "). 

% 2. scan/2
test(scan) :-
    scan("hello world ", Words),
    assertion(Words == ["hello","world"]).

% 3. remove_stop_words/3
test(remove_stop_words) :-
    remove_stop_words(
      ["this","is","a","test"], ["is","a"], Filtered),
    assertion(Filtered == ["this","test"]).

% 4. frequencies/2
test(frequencies) :-
    frequencies(["a","b","a"], Freq),
    assertion(Freq == [a-2, b-1]).

% 5. sorted/2
test(sorted) :-
    sorted([b-1, a-2], Sorted),
    assertion(Sorted == [a-2, b-1]).

% 6. read_stop_words/2 (用暫存檔模擬)
test(read_stop_words, [
    setup((
      open('tmp_sw.txt', write, S1),
      write(S1, "a,is,the"), close(S1)
    )),
    cleanup(delete_file('tmp_sw.txt'))
]) :-
    read_stop_words('tmp_sw.txt', SW),
    % 前三個應該是原始三個停用詞
    nth0(0, SW, "a"),
    nth0(1, SW, "is"),
    nth0(2, SW, "the"),
    % 一共 3 + 26 個字
    length(SW, Len),
    Len == 29.

% 7. 結合測試：word_frequencies/2 會印出「test - 2」和「only - 1」
test(word_frequencies_output, [
    setup((
      open('in.txt', write, InS),
      write(InS, "This is a test, only a test."), close(InS),
      open('sw.txt', write, SwS),
      write(SwS, "a,is,the"), close(SwS)
    )),
    cleanup((delete_file('in.txt'), delete_file('sw.txt')))
]) :-
    with_output_to(string(Out),
        word_frequencies('in.txt','sw.txt')),
    assertion(sub_string(Out, _, _, _, "test - 2")),
    assertion(sub_string(Out, _, _, _, "only - 1")).

:- end_tests(word_freq).
