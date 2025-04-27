import tensorflow as tf
import tensorflow_hub as hub

class MobileNetV2Model(tf.keras.Model):
    def __init__(self):
        super(MobileNetV2Model, self).__init__()
        self.preprocessing = tf.keras.layers.Resizing(224, 224)
        self.hub_layer = hub.KerasLayer(
            "https://tfhub.dev/google/tf2-preview/mobilenet_v2/classification/4",
            trainable=False
        )

    def call(self, inputs):
        x = self.preprocessing(inputs)
        return self.hub_layer(x)

# Create and compile the model
model = MobileNetV2Model()
# Build the model with a sample input
model.build((None, None, None, 3))

# Save with signature
input_signature = tf.TensorSpec([None, None, None, 3], tf.float32)
@tf.function(input_signature=[input_signature])
def serving_fn(x):
    return {'outputs': model(x)}

tf.saved_model.save(model, "saved_mobilenet_v2", signatures={"serving_default": serving_fn})
