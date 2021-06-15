<html>
    <head>
        <title>
            index
        </title>
        <body>
            <form action="" method="post">
                <input type="text" name="param1" size="20" maxlength="20" value= "{{ .Value1 }}" />
                
                <select id="type" name="param3">
                    <option value="1" {{ if eq .Value3 "1" }} selected {{ end }}>+</option>
                    <option value="2" {{ if eq .Value3 "2" }} selected {{ end }}>-</option>
                    <option value="3" {{ if eq .Value3 "3" }} selected {{ end }}>×</option>
                    <option value="4" {{ if eq .Value3 "4" }} selected {{ end }}>÷</option>
                </select>

                <input type="text" name="param2" size="20" maxlength="20" value="{{ .Value2 }}" />
                <p>Result {{ .Result }}</p>
                <input type="submit" name="送信" value="Submit" />
            </form>
        </body>
    </head>
</html>