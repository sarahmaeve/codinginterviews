##
# justify(13, "This is a good sentence to try.")
#
#  Might return:
# This   is   a
# good sentence
# to try.
#
# return value: "This  is   a\ngoood sentence\nto try."
##

## good
## sentences
## supercalifragilistic
## dealing with punctuation on start of line

# split into words (words including punctuation)
# word : non-space
# counter per line (loop by character)
# breakup : adding \n
# split and reassemble?
# return : single string
# where to add \b?
def reassemble(line_length, word_list):
    # try to distribute equitably
    # a .. b .. c . d
    # != a . b . c ... d
    word_length = len("".join(word_list))
    spaces_req = line_length - word_length
    locations = len(word_list) - 1
    divided_spaces = int(spaces_req/locations)
    remainder = spaces_req % locations
    output = ""
    for word in word_list:
        output += word + "_" * divided_spaces
        if remainder > 0:
            output += "_"
            remainder -= 1

    return output + '\n'

def justify(line_length, candidate_string):
    # or match
    current_length = 0
    start_slice = 0
    reformatted_string = ""
    candidate_arr = candidate_string.split()
    for i, word in enumerate(candidate_arr):
        current_length += len(word)
        min_padding = i - start_slice # what happens with a single word on a line, [0] [1] do I return -1
        if current_length + min_padding == line_length:
            # join
            line = reassemble(line_length, candidate_arr[start_slice:i+1])
            reformatted_string += line
            start_slice = i+1
            current_length = 0
        elif current_length + min_padding < line_length:
            current_length
            continue # next?
        elif current_length + min_padding > line_length:
            line = reassemble(line_length, candidate_arr[start_slice:i])
            reformatted_string += line
            start_slice = i
            current_length = len(word)
    else:
        # line = reassemble(line_length, candidate_arr[start_slice:])
        line = " ".join(candidate_arr[start_slice:])
        reformatted_string += line + "\n"
    # missing case: word > line_length
    return reformatted_string


print(justify(13, "This is a good sentence to try."))




