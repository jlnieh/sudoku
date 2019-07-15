
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
                    htmlPuzzle += '<td><input type="text" class="square" maxlength="1" size="1" id="s' + id + '" tabindex="' + (id+1).toString() + '" value=""/></td>\n';
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

function splitintosquares(grid) {
    var s = 0;
    for (var i=0; i<grid.length; i++) {
        var c = grid.charAt(i);
        if ((c >= '1') && (c <= '9')) {
            $("#s" + s).val(c);
            s++;
        }
        else if ((c=='0') || (c == '.')) {
            $("#s" + s).val('');
            s++;
        }
        else {
            continue;
        }
        if (s >= 81) break;
    }
    while (s < 81) {
        $("#s" + s).val('');
        s++;
    }
}

function generatevalues(i, v) {
    var grid = $("#grid").val();
    if (grid.length < 81) {
        grid += ".".repeat(81 - grid.length)
    }
    if ((v > 0) && (v <= 9)) {
        grid = grid.substr(0, i) + v + grid.substr(i+1);
    }
    else {
        grid = grid.substr(0, i) + '.' + grid.substr(i+1);
    }
    $("#grid").val(grid);
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
                if ((data.solved) && (data.values) && (data.values.length == 81)) {
                    drawSolution(data.values);
                } else {
                    alert('Cannot solve the puzzle!')
                }
            },
            error: function(xhr, error){
                alert("Something wrong to communicate with solver!");
                console.debug(xhr); 
                console.debug(error);
            }
        });
        e.preventDefault();
    });

    $("#grid").on('keyup change', function(){
        $("#solution").html('');
        splitintosquares($(this).val());
    });

    $(".square").on('keyup change', function(){
        var id = parseInt($(this).attr('id').substr(1));
        var vl = $(this).val();

        $("#solution").html('');
        generatevalues(id, vl);
        if (vl.length == 1) {
            if (id < 80) {
                $("#s" + (id+1).toString()).focus().select();
            }
            else {
                $("#grid").focus();
            }
        }
    });
});
