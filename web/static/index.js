
function drawEmptyPuzzle(){
    var htmlPuzzle = '<table style="background-color:black;">';
    for (var R = 0; R < 3; R++){
        htmlPuzzle += "<tr>\n";
        for (var C = 0; C < 3; C++){
            htmlPuzzle += "<td>";
            htmlPuzzle += '<table style="background-color:gray;">';
            for (var r = 0; r < 3; r++){
                htmlPuzzle += "<tr>\n";
                for (var c = 0; c < 3; c++){
                    var id = (R*3 + r)*9 + (C*3 + c);
                    htmlPuzzle += '<td><input type="text" class="square" maxlength="1" size="1" id="s' + id + '" value=""/></td>\n';
                }
                htmlPuzzle += "</tr>\n";
            }
            htmlPuzzle += '</table>';
            htmlPuzzle += "</td>";
        }
        htmlPuzzle += "</tr>\n";
    }
    htmlPuzzle += '</table>';

    $("#puzzle").html(htmlPuzzle);
}

function drawSolution(data){
    var htmlPuzzle = '<table style="background-color:black;">';
    for (var R = 0; R < 3; R++){
        htmlPuzzle += "<tr>\n";
        for (var C = 0; C < 3; C++){
            htmlPuzzle += "<td>";
            htmlPuzzle += '<table style="background-color:gray;">';
            for (var r = 0; r < 3; r++){
                htmlPuzzle += "<tr>\n";
                for (var c = 0; c < 3; c++){
                    var id = (R*3 + r)*9 + (C*3 + c);
                    htmlPuzzle += '<td>' + data[id] + '</td>\n';
                }
                htmlPuzzle += "</tr>\n";
            }
            htmlPuzzle += '</table>';
            htmlPuzzle += "</td>";
        }
        htmlPuzzle += "</tr>\n";
    }
    htmlPuzzle += '</table>';

    $("#solution").html(htmlPuzzle);
}

$(document).ready(function(){
    drawEmptyPuzzle();

    $("#target").submit(function(e){
        var form = $(this);
        var url = form.attr('action');

        $.ajax({
            type: "POST",
            url: url,
            data: form.serialize(),
            success: function(data)
            {
                drawSolution(data);
            },
            error: function(xhr, error){
                alert("Cannot solve the puzzle!");
                console.debug(xhr); 
                console.debug(error);
            }
        });
        e.preventDefault();
    });

    $("#grid").change(function(e){
        $("#solution").html('');
    });

});
