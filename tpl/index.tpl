<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Docker Captain's Atom Updates</title>
  <link rel='alternate' type='application/atom+xml' title='Atom 0.3' href='http://argh.gianarb.it/index.xml'>
  <style>
    .logo img {
        max-height: 300px;
    }
  </style>
</head>
<body>
    <div class="logo">
        <img src="/captains-logo.png" alt="Docker Captains logo">
    </div>

    <div class="slogan">
        <h1>Arghs and news from Docker Captains</h1>
    </div>

    <div class="content">
        <table>
            {{ range .Items }}
            <tr>
                <td><a href="{{ .Link.Href }}" target="_blank">{{ .Title }}</a></td>
            </tr>
            {{ end }}
        </table>

    </div>

<script>
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

ga('create', 'UA-83613100-1', 'auto');
ga('send', 'pageview');

</script>
</body>
</html>
