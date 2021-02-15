import sys

def create_word_dict(filename, maxlen=15):
    word_termination_key = "."
    word_termination_value = "."
    with open(filename) as file:
        word_array = []
        for line in file:
            line = line.rstrip().lower()
            word_array.append(line)
        
    word_dict = {}
    total_length = len(word_array) + 1
    word_count = 1
    for word in word_array:
        sys.stdout.write(f"{word_count}/{total_length}\r")
        position = word_dict
        last_letter = None
        i = 0
        for letter in word:
            if i >= maxlen:
                break
            
            try:
                position = position[letter]
            except KeyError:
                position[letter] = {}
                position = position[letter]
                
            last_letter = letter
            i += 1
        position[word_termination_key] = word_termination_value
        word_count += 1
    sys.stdout.write(f"Finished converting {total_length} words\n")

    return word_dict

def write_json_dict(word_dict, filename):
    import json
    with open(filename, 'w') as file:
        json.dump(word_dict, file)
    sys.stdout.write("Saved output in " + filename)

def load_json(filename):
    import json
    with open(filename, 'rb') as file:
        return json.loads(file.read())

def check_words(word_dict, string):
    def verify_dict(sub_dict, string):
        for entry in sub_dict:
            if entry == ".":
                print(string)
            else:
                verify_dict(sub_dict[entry], string + entry)
    
    dict_position = word_dict
    for letter in string:
        dict_position = dict_position[letter]
    
    verify_dict(dict_position, string)


def main(argv):
    filename_input = argv[0]
    print("Converting " + filename_input + "...")
    word_dict = create_word_dict(filename_input)
    filename_output = "".join(filename_input.split(".")[:-1]) + "_graph.json"
    write_json_dict(word_dict, filename_output)
    sys.stdout.write("\n")

if __name__ == "__main__":
   main(sys.argv[1:])
