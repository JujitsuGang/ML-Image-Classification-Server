
# ML-Image-Classification-Server
A server-side application for image classification using K-Nearest Neighbours (KNN) algorithm


- Reads all the image datasets from the /dataset directory
- Allows for image upload for classification
- Accepts PNG, JPG, and JPEG files
- All uploaded images are resized to a standard size and converted to PNG as configured in config.json
- Calculated euclidean distances for input images
- Applies KNN (K-Nearest Neighbours algorithm) to classify images
- The server then returns the result of the classification, i.e., the label of the object in the image


## Instructions

Put the image dataset in /dataset directory. Each subdirectory should contain images of a single class or category as shown below:

```
dataset/
    leopard/
        img01.png
        img02.png
        img03.png
        ...
    laptop/
        img01.png
        img02.png
        ...
    camera/
        img01.png
        img02.png
        ...
```

To run the server:
```
    >./ML-Image-Classification-Server
```

Next, perform requests with new images to get a response from the server classifying them:

```
curl -F file=@./testimage.png http://127.0.0.1:3055/image
```

The server will return a response like:
```
    seems to be a leopard
```

You can perform tests using the test.sh file:
```
bash test.sh
```

## Useful commands

To send files over ssh:
```
scp dataset.tar.gz root@SERVERIP:/root/ML-Image-Classification-Server
```

To untarr files on the server:
```
tar -xvzf dataset.tar.gz
```