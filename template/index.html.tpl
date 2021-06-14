<html>
    <head>
        <title>
            index
        </title>
        <body>
            <form action="" method="post">
                <input type="text" name="param1" size="20" maxlength="20" value= "{{ .Value1 }}" />
                <input type="text" name="param2" size="20" maxlength="20" value="{{ .Value2 }}" />
                <p>Result {{ .Result }}</p>
                <input type="submit" name="送信" value="Submit" />
            </form>
        </body>
    </head>
</html>