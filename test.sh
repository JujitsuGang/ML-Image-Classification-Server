
echo "sending leopard to server"
echo "server response:"
curl -F file=@./testleopard1.png http://127.0.0.1:3055/image
echo ""

echo "sending leopard to server"
echo "server response:"
curl -F file=@./testleopard2.png http://127.0.0.1:3055/image
echo ""