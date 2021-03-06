function submitForm(){
    let inputData = {};
    $(":input").each(function() {
        inputData[$(this).attr("name")] = parseInt($(this).val());
    });
    if (inputData.numberOfPlayers % inputData.numberOfPlayersPerLine == 0) {
        $("p").html("Sorry, you don't need help as you number of players can be divided by the number of players per line");
        return
    }
    if (inputData.numberOfPlayers < 2 * inputData.numberOfPlayersPerLine) {
        $("p").html("Sorry, you don't have enough players to build 2 lines");
        return
    }
    $.ajax({
        type: "get",
        url: (window.location.pathname.startsWith('/web/hockeyKidsLinesPage.html') ? '../api' : '/api') + '?numberOfPlayers=' + inputData.numberOfPlayers + '&numberOfPlayersPerLine=' + inputData.numberOfPlayersPerLine + '&numberOfLinesPerMatch=' + inputData.numberOfLinesPerMatch,
        success: displayResult,
        error: function (data) {
            $("p").html("Sorry, an error occurred during the processing of your request, please try again later");
            $("form").remove()
        },
    });
}

function displayResult(data) {
    $("p").html("");
    let root = $(result);
    let innerDiv = "";
    let bestMatch = JSON.parse(data).bestMatch;
    if (bestMatch.length == 0) {
        $("p").html("Sorry, we didn't manage to find a correct solution to your problem");
    } else {
        for (let i of bestMatch) {
            innerDiv += buildMatchView(i) + "\n"
        }
        root.html(innerDiv)
    }
}

function buildMatchView(data) {
    let playerNames = $("select").val();
    for (let i = playerNames.length + 1; i <= data.length; i++) {
        playerNames.push(`Player ${i}`);
    }
    let result = ["<table border='1'>"];
    let i = 0;
    result.push(`<tr><th>Line Number</th>`);
    for (let firstRow of data[0]) {
        result.push(`<th>${++i}</th>`);
    }
    result.push(`<th>Total</th>`);
    i = 0;
    result.push("</tr>");
    for(let row of data) {
        result.push(`<tr><th>${playerNames[i++]}</th>`);
        let cpt = 0;
        for(let cell of row.split("").reverse().join("")){
            if (cell != 0) {
                cpt++;
            }
            result.push(`<td>${cell == '0' ? '': 'X'}</td>`);
        }
        result.push(`<td align="center">${cpt}</td>`);
        result.push("</tr>");
    }
    result.push("</table><br>");
    return result.join('\n');
}