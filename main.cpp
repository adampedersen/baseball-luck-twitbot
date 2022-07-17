#include <iostream>
#include <iomanip>
#include <fstream>

using namespace std;

//script for generating urls from player names and ids
int main(int argc, char* argv[]){
	// file pointer
	std::ofstream fout;
    std::ifstream fin;

	fin.open("fangraphs-id-map.csv");
    // opens an existing csv file or creates a new file.
    fout.open("fangraphs-url-map.csv", ios::out | ios::app);

    string lineOld;
    string lineNew;
    string playerName;
    string url;
    int idxOfComma;//(of comma)
    int idxOfSpace;

	while(getline(fin,lineOld)){
        idxOfComma = lineOld.find(',');
        playerName = lineOld.substr(0, idxOfComma);
        lineNew = lineOld;
        while(lineNew.find(' ') != std::string::npos){
            idxOfSpace = lineNew.find(' ');
            lineNew[idxOfSpace] = '-';
        }
        lineNew[idxOfComma] = '/';
        lineNew = playerName + ","+ "https://www.fangraphs.com/players/" + lineNew + "\n";
        fout<<lineNew;
    }
    fout.close();
    fin.close();


}