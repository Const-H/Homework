#include <iostream>
#include <fstream>
#include <vector>

using namespace std;

vector<vector<double>> loadMatrixFromFile() {
    ifstream inputFile("matrix.txt");
    vector<vector<double>> data;

    if (!inputFile) {
        cout << "Ошибка открытия файла!" << endl;
        return data;
    }

    double value;
    while (inputFile >> value) {
        vector<double> row;
        row.push_back(value);

        while (inputFile.peek() != '\n' && inputFile.peek() != EOF) {
            inputFile >> value;
            row.push_back(value);
        }

        data.push_back(row);
    }

    inputFile.close();
    return data;
}

double computeDiagonalSum(const vector<vector<double>>& m) {
    double sum = 0;
    for (int i = 0; i < m.size(); i++) {
        sum += m[i][i];
    }
    return sum;
}

vector<vector<double>> createTransposed(const vector<vector<double>>& m) {
    vector<vector<double>> result(m[0].size(), vector<double>(m.size()));

    for (int i = 0; i < m.size(); i++) {
        for (int j = 0; j < m[0].size(); j++) {
            result[j][i] = m[i][j];
        }
    }

    return result;
}

double computeMatrixDeterminant(vector<vector<double>> m) {
    int size = m.size();

    if (size == 1) return m[0][0];
    if (size == 2) {
        return m[0][0] * m[1][1] - m[0][1] * m[1][0];
    }

    double determinant = 0;

    for (int col = 0; col < size; col++) {
        vector<vector<double>> submatrix(size - 1, vector<double>(size - 1));

        for (int i = 1; i < size; i++) {
            int subcol = 0;
            for (int j = 0; j < size; j++) {
                if (j == col) continue;
                submatrix[i - 1][subcol] = m[i][j];
                subcol++;
            }
        }

        double sign = (col % 2 == 0) ? 1 : -1;
        determinant += sign * m[0][col] * computeMatrixDeterminant(submatrix);
    }

    return determinant;
}

void saveResults(double diagSum, double det, const vector<vector<double>>& transposed) {
    ofstream outputFile("results.txt");

    outputFile << "Сумма диагонали: " << diagSum << endl;
    outputFile << "Определитель: " << det << endl;
    outputFile << "Транспонированная матрица:" << endl;

    for (const auto& row : transposed) {
        for (double num : row) {
            outputFile << num << " ";
        }
        outputFile << endl;
    }

    outputFile.close();
}

int main() {
    vector<vector<double>> matrixData = loadMatrixFromFile();

    if (matrixData.empty() || matrixData.size() != matrixData[0].size()) {
        cout << "Ошибка: матрица не квадратная" << endl;
        return 1;
    }

    double diagSum = computeDiagonalSum(matrixData);
    double det = computeMatrixDeterminant(matrixData);
    vector<vector<double>> transposedMatrix = createTransposed(matrixData);

    saveResults(diagSum, det, transposedMatrix);

    cout << "Результаты сохранены в results.txt" << endl;

    return 0;
}