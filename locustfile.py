import time
import random
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):

    # @task
    # def view_items(self):
    #     for item_id in self.uuids:
    #         self.client.get(f"/key/{item_id}", name="/item")
    #         time.sleep(0.1)
    @task
    def view_items(self):
        random_index = random.randint(0,len(self.uuids)-1)
        self.client.get(f"/key/{self.uuids[random_index]}", name="/item")

    def on_start(self):
        with open("500ids") as f:
            content_list = f.readlines()
        content_list = [x.strip() for x in content_list]
        self.uuids=content_list