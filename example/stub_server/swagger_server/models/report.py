# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.amount import Amount  # noqa: F401,E501
from swagger_server.models.area import Area  # noqa: F401,E501
from swagger_server import util


class Report(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """
    def __init__(self, area: Area=None, amount: Amount=None):  # noqa: E501
        """Report - a model defined in Swagger

        :param area: The area of this Report.  # noqa: E501
        :type area: Area
        :param amount: The amount of this Report.  # noqa: E501
        :type amount: Amount
        """
        self.swagger_types = {
            'area': Area,
            'amount': Amount
        }

        self.attribute_map = {
            'area': 'area',
            'amount': 'amount'
        }
        self._area = area
        self._amount = amount

    @classmethod
    def from_dict(cls, dikt) -> 'Report':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The report of this Report.  # noqa: E501
        :rtype: Report
        """
        return util.deserialize_model(dikt, cls)

    @property
    def area(self) -> Area:
        """Gets the area of this Report.


        :return: The area of this Report.
        :rtype: Area
        """
        return self._area

    @area.setter
    def area(self, area: Area):
        """Sets the area of this Report.


        :param area: The area of this Report.
        :type area: Area
        """
        if area is None:
            raise ValueError("Invalid value for `area`, must not be `None`")  # noqa: E501

        self._area = area

    @property
    def amount(self) -> Amount:
        """Gets the amount of this Report.


        :return: The amount of this Report.
        :rtype: Amount
        """
        return self._amount

    @amount.setter
    def amount(self, amount: Amount):
        """Sets the amount of this Report.


        :param amount: The amount of this Report.
        :type amount: Amount
        """
        if amount is None:
            raise ValueError("Invalid value for `amount`, must not be `None`")  # noqa: E501

        self._amount = amount