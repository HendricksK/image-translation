# image translation service

translates text files that houses image links to json files, based on topic. eg: I want all images related to topic "the+room+memes"
text file would just be named after topic, we search for filename named after topic, consume text inside it, will basically just be CSV
and return all the strings as neat json

[
    {
        "requestedTopic": "the+room+memes",
        "Images": [
            "https://preview.redd.it/0ndypduzlf641.jpg?width=640\u0026crop=smart\u0026auto=webp\u0026s=5cf49a2deb07cfa5bf628e1dfb074b0452bb8041",
            "https://preview.redd.it/q6ez8ub45i931.jpg?width=640\u0026crop=smart\u0026auto=webp\u0026s=4b4859c1cc752ac3310496852b8fd19a77aa7fc7"
        ],
        "Number": 2
    }
]
