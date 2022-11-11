export function safeAlphabetString(string) {
    let regex = new RegExp(".?[^A-z ]+.?", "g");
    if (regex.test(string) || string == ""){
        return false;
    }   
    return true;
}