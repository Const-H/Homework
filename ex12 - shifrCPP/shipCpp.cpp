#include <iostream>
#include <fstream>
#include <string>
#include <cctype>

using namespace std;

string reverseAlphabet(const string& text) {
    string output;
    for (char c : text) {
        if (isupper(c)) {
            output += 'Z' - (c - 'A');
        } else if (islower(c)) {
            output += 'z' - (c - 'a');
        } else {
            output += c;
        }
    }
    return output;
}

string shiftLetters(const string& text, int key) {
    string output;
    for (char c : text) {
        if (isupper(c)) {
            output += ((c - 'A' + key) % 26) + 'A';
        } else if (islower(c)) {
            output += ((c - 'a' + key) % 26) + 'a';
        } else {
            output += c;
        }
    }
    return output;
}

int main() {
    ifstream input("input.txt");
    if (!input) {
        cout << "Файл input.txt не найден" << endl;
        return 1;
    }

    string data;
    getline(input, data);
    input.close();

    int key;
    cout << "Укажите ключ сдвига: ";
    cin >> key;

    string reversed = reverseAlphabet(data);
    string shifted = shiftLetters(data, key);

    ofstream output("output.txt");
    if (!output) {
        cout << "Ошибка создания output.txt" << endl;
        return 1;
    }

    output << "Реверсированный алфавит: " << reversed << endl;
    output << "Сдвинутые буквы (ключ " << key << "): " << shifted << endl;
    output.close();

    cout << "Операция завершена. Проверьте output.txt" << endl;

    return 0;
}